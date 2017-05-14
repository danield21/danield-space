const Anime = require('animejs')
const Balloons = require('./balloons')
const Stick = require('./stick')
const util = require('./util')
const router = require('./router')
const Sun = require('./sun')
const Modernizr = require('modernizr')
const Please = require('pleasejs')
const sdRand = require('gauss-random')

function meetsRequirements() {
    return Modernizr.eventlistener &&
        Modernizr.queryselector &&
        Modernizr.es5 &&
        Modernizr.promises
}

function styleDesert(mainCloud, desert) {
    mainCloud.style.marginBottom = 0
    desert.style.visibility = 'hidden'
    return new Promise((resolve) => {
        setTimeout(() => {
            desert.style.visibility = 'visible'
            let screen = util.screenSize()
            mainCloud.style.marginBottom = (screen.height - 300) + 'px'
            resolve()
        }, 0)
    })
}

document.addEventListener('DOMContentLoaded', () => {
    if (!meetsRequirements()) {
        return
    }

    const main = Bliss('main')
    const easel = Bliss('#sky-easel')

    initEasel(easel)

    if (Balloons.meetsRequirements()) {
        initBalloons(easel)
    }

    initSun(easel)

    if (router.meetsRequirements()) {
        initRouting(main)
    }

    let mainCloud = Bliss('body > .cloud')
    let desert = Bliss('body > footer > .sand')
    let styleFunc = styleDesert.bind(null, mainCloud, desert)

    let mountainRange = document.getElementById('mountain-range')
    let stickFunc = Stick.toBottom(mountainRange)
    let raiseEasel = () => {
        var height = util.screenSize().height
        const offset = Bliss('.sand').getBoundingClientRect().top
        if (offset < height) {
            easel.style.transform = `translateY(-${height - offset}px)`
        } else {
            easel.style.transform = null
        }
    }

    Bliss.$('.js-fillNow').forEach(button => {
        button.addEventListener('click', e => {
            var id = e.target.dataset.target
            var input = document.getElementById(id)
            util.setDateTimeInputToNow(input)
        })
    })

    styleFunc()
        .then(stickFunc)
        .then(raiseEasel)
    window.addEventListener('load', () => {
        styleFunc()
            .then(stickFunc)
            .then(raiseEasel)
    })
    window.addEventListener('resize', () => {
        styleFunc()
            .then(stickFunc)
            .then(raiseEasel)
    })
    document.addEventListener('scroll', () => {
        stickFunc().then(raiseEasel)
    })
})

function initEasel(easel) {
    easel.style.position = 'fixed'
    easel.style.top = 0
}

function initSun(easel) {
    const sun = Sun.create(1000)
    Sun.setColor(sun, '#FFF250', '#FFFFFF')
    easel.appendChild(sun)
}


const balloons = {
    MAX_AMOUNT: 10,
    EVERY: 20000,
    MIN_HEIGHT: 100,
    AVG_HEIGHT: 150,
    MAX_HEIGHT: 200
}

function initBalloons(easel) {
    Bliss.fetch('/dist/svg/balloon.svg').then(svg => {
        const intID = setInterval(addBalloon, balloons.EVERY)
        addBalloon()

        function addBalloon() {
            if (easel.childNodes.length >= balloons.MAX_AMOUNT || (document.hidden || document.msHidden || document.webkitHidden)) {
                return
            }
            const screen = util.screenSize()
            const hHalf = screen.height / 2
            const top = util.choosePoint(sdRand(), 0, hHalf, screen.height - balloons.MAX_HEIGHT)
            const left = screen.width
            const hexColor = Please.make_color({
                saturation: .8 + Math.random() * .2,
                value: .8 + Math.random() * .2
            })[0]
            const height = util.choosePoint(sdRand(), balloons.MIN_HEIGHT, balloons.AVG_HEIGHT, balloons.MAX_HEIGHT)
            Balloons.parseSVG(svg)
                .then(Balloons.size(height))
                .then(Balloons.position(top, left))
                .then(Balloons.storeIds)
                .then(Balloons.color(hexColor))
                .then(Balloons.drawOn(easel))
                .then(Balloons.fly)
                .then(Balloons.remove(easel))
                .catch(() => clearInterval(intID))
        }
    })
}

function initRouting(main) {
    router.init()

    const outFunc = transitionOut(main)
    const inFunc = transitionIn(main)

    window.addEventListener('click', e => {
        router.handleRouting(outFunc, inFunc)(e)
    })

    window.addEventListener('submit', e => {
        router.handleForm(outFunc, inFunc)(e)
    })
    window.addEventListener('popstate', e => {
        main.innerHTML = e.state
        window.dispatchEvent(new Event('resize'))
    })
}

const transitionChildrenClass = 'transition-children'

function transitionTarget(elem) {
    if (!elem.classList.contains(transitionChildrenClass) || elem.children == null || elem.children.length == 0) {
        return [elem]
    }
    return Array.from(elem.children).reduce((list, e) => list.concat(transitionTarget(e)), [])
}

function transitionOut(main) {
    return () => {
        let children = Array.from(main.children)
        let targets = children.reduce((list, e) => list.concat(transitionTarget(e)), [])
        scrollToTop(250)
        const a = Anime({
            targets: targets,
            duration: 500,
            easing: 'linear',
            translateY: (_, i) => `-${(i+1) * 100}px`,
            opacity: 0
        })
        return a.finished.then(() => {
            children.forEach(child => main.removeChild(child))
            window.dispatchEvent(new Event('resize'))
        })
    }
}

function scrollToTop(scrollDuration) {
    const scrollStep = -window.scrollY / (scrollDuration / 15)
    const scrollInterval = setInterval(function() {
        if (window.scrollY != 0) {
            window.scrollBy(0, scrollStep)
        } else clearInterval(scrollInterval)
    }, 15)
}

function transitionIn(main) {
    return ([_, frag]) => {
        let targets = Array.from(frag.children).reduce((list, e) => list.concat(transitionTarget(e)), [])
        targets.forEach((c, i) => {
            c.style.transform = `translateY(-${(i+1) * 100}px)`
            c.style.opacity = '0'
        })
        const a = Anime({
            targets: targets,
            duration: 500,
            easing: 'linear',
            translateY: '0',
            opacity: 1
        })
        main.appendChild(frag)
        return a.finished
    }
}