const Modernizr = require('modernizr')
const util = require('./util')

module.exports = {
    init,
    next,
    meetsRequirements,
    handleRouting,
    handleForm,
    handleBack
}

function meetsRequirements() {
    return Modernizr.documentfragment &&
        Modernizr.xhrresponsetypedocument &&
        Modernizr.history
}

function init() {
    const page = window.location.href
    storeState(Bliss('main'), {
        url: page
    })
}

function isExternalLink(a) {
    return a.hostname.length && window.location.hostname !== a.hostname
}

function dontLoadAjax(e) {
    return e.classList.contains('router-no-spa')
}

function storeState(main, state, newPage, scroll) {
    state.id = getUniqueId(window.localStorage)

    if (window.history.state && scroll) {
        window.history.state.scroll = scroll
        window.history.replaceState(window.history.state, window.document.title, window.location.href)
    }

    if (newPage) {
        window.history.pushState(state, window.document.title, state.url)
    } else {
        window.history.replaceState(state, window.document.title, state.url)
    }

    window.localStorage.setItem(state.id, main.innerHTML)
}

function next(url, scroll) {
    return html => {
        const main = html.body

        storeState(main, {
            url,
        }, window.location.href !== url, scroll)

        const frag = document.createDocumentFragment()
        Array.from(main.children).forEach(c => frag.appendChild(c))
        return Promise.resolve(frag)
    }
}

const prefix = 'page-'

function getUniqueId(storage) {
    let id = guid()

    while (storage.getItem(prefix + id)) {
        id = guid()
    }
    return prefix + id
}

function guid() {
    return s4() + '-' + s4() + '-' + s4() + '-' + s4()
}

function s4() {
    return Math.floor((1 + Math.random()) * 0x10000)
        .toString(16)
        .substring(1)
}

function submitForm(form) {
    var url = form.getAttribute('action')
    var method = form.method
    var data = Array.from(form.elements).reduce((encode, e) => {
        if (e.name) {
            if (encode) encode += '&'
            return encode + encodeURIComponent(e.name) + '=' + encodeURIComponent(e.value)
        }
        return encode
    }, '')

    return getPage(url, {
        method,
        data,
        headers: {
            Accept: 'application/json'
        }
    })
}

function navigate(a) {
    if (a && a.tagName.toUpperCase() != 'A' && a.href) {
        return Promise.reject(new Error('Provided value is not an A element with an href'))
    }

    return getPage(a.href, {
        headers: {
            Accept: 'application/json'
        }
    })
}

function getPage(url, data) {
    return Bliss.fetch(url + (url.indexOf('?') >= 0 ? '&' : '?') + 'content-format=json', data).then(
        response => new Promise((resolve, _reject) => {
            resolve(JSON.parse(response.response))
        }), e => new Promise((resolve, _reject) => {
            resolve(JSON.parse(e.xhr.response))
        })
    ).then(page => {
        const parser = new DOMParser()
        page.Content = parser.parseFromString(page.Content, 'text/html')
        return Promise.resolve(page.Content)
    })
}

function handleBack(main, transitionOut, transitionIn) {
    return e => {
        if (e.defaultPrevented) {
            return
        }

        transitionOut().then(() => {
            main.innerHTML = window.localStorage.getItem(e.state.id)
            return Promise.resolve(e.state)
        }).then(transitionIn)
    }
}

function handleRouting(transitionOut, transitionIn) {
    return e => {
        if (e.defaultPrevented) {
            return
        }

        let a = util.findAncestor(e.target, e => e.tagName.toUpperCase() == 'A')

        if (a == null || isExternalLink(a) || dontLoadAjax(a)) {
            return
        }

        e.preventDefault()

        Promise.all([
            transitionOut(),
            navigate(a).then(next(a.href, { y: window.scrollY, x: window.scrollX }))
        ]).then(transitionIn)
    }
}

function handleForm(transitionOut, transitionIn) {
    return e => {
        if (e.defaultPrevented) {
            return
        }

        let form = util.findAncestor(e.target, e => e.tagName.toUpperCase() == 'FORM')

        if (form == null) {
            return
        }

        e.preventDefault()

        Promise.all([
            transitionOut(),
            submitForm(form).then(next(form.action, { y: window.scrollY, x: window.scrollX }))
        ]).then(transitionIn)
    }
}