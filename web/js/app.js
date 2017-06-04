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

    document.addEventListener('click', e => {
        const button = util.findAncestor(e.target, element => element.classList.contains('js-fillNow'))
        if (button != null) {
            const id = button.dataset.target
            const input = document.getElementById(id)
            util.setDateTimeInputToNow(input)
        }
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
            if (Bliss.$('.svg-balloon', easel).length >= balloons.MAX_AMOUNT || (document.hidden || document.msHidden || document.webkitHidden)) {
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

    const handleForm = router.handleForm(outFunc, inFunc)
    const handleRouting = router.handleRouting(outFunc, inFunc)
    const handleBack = router.handleBack(main, () => Promise.resolve(), (state) => {
        window.dispatchEvent(new Event('resize'))
        if (state.scroll) {
            setTimeout(() => {
                window.scrollTo(state.scroll.x, state.scroll.y)
            }, 0)
        }
        return Promise.resolve()
    })

    window.addEventListener('click', handleRouting)

    window.addEventListener('submit', handleForm)
    window.addEventListener('popstate', handleBack)
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
        util.scrollToTop(250)
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
        return a.finished.then(() => {
            window.dispatchEvent(new Event('resize'))
        })
    }
}