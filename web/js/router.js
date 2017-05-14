const Modernizr = require('modernizr')
const util = require('./util')

module.exports = {
    init,
    next,
    meetsRequirements,
    handleRouting,
    handleForm,
}

function meetsRequirements() {
    return Modernizr.documentfragment &&
        Modernizr.xhrresponsetypedocument &&
        Modernizr.history
}

function init() {
    const page = window.location.origin + window.location.pathname
    window.history.replaceState(Bliss('main').innerHTML, window.document.title, page)
}

function isExternalLink(a) {
    return a.hostname.length && window.location.hostname !== a.hostname
}

function dontLoadAjax(e) {
    return e.classList.contains('router-no-spa')
}

function next(url, perform) {
    return perform.then(html => {
        const main = Bliss('main', html)
        if (window.location.href !== url) {
            window.history.pushState(main.innerHTML, window.document.title, url)
        } else {
            window.history.replaceState(main.innerHTML, window.document.title, url)
        }
        const frag = document.createDocumentFragment()
        Array.from(main.children).forEach(c => frag.appendChild(c))
        return Promise.resolve(frag)
    }, e => {
        return Promise.reject(e)
    })
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

    return Bliss.fetch(url, { responseType: 'document', method, data }).then(response => {
        return response.responseXML ? Promise.resolve(response.responseXML) : Promise.reject(new Error('Did not get a document back from ' + url + '?' + data))
    }, e => {
        return e.xhr.responseXML ? Promise.resolve(e.xhr.responseXML) : Promise.reject(new Error('Did not get a document back from ' + url + '?' + data))
    })
}

function navigate(a) {
    if (a && a.tagName.toUpperCase() != 'A' && a.href) {
        return Promise.reject(new Error('Provided value is not an A element with an href'))
    }
    const url = a.href

    return Bliss.fetch(a.href, { responseType: 'document' }).then(response => {
        return response.responseXML ? Promise.resolve(response.responseXML) : Promise.reject(new Error('Did not get a document back from ' + url))
    }, e => {
        return e.xhr.responseXML ? Promise.resolve(e.xhr.responseXML) : Promise.reject(new Error('Did not get a document back from ' + url))
    })
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
            next(a.href, navigate(a))
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
            next(form.action, submitForm(form))
        ]).then(transitionIn)
    }
}