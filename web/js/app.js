import Anime from 'animejs'
import * as Modernizr from 'modernizr'
import * as Please from 'pleasejs'
import gaussRandom from 'gauss-random'

import * as util from './util'

import * as Balloons from './balloons'
import * as Router from './router'
import * as Sun from './sun'
import * as desert from './desert'
import * as Easel from './easel'

function meetsRequirements() {
    return Modernizr.eventlistener &&
        Modernizr.queryselector &&
        Modernizr.es5array &&
        Modernizr.es5date &&
        Modernizr.es5function &&
        Modernizr.es5object &&
        Modernizr.strictmode &&
        Modernizr.es5string &&
        Modernizr.json &&
        Modernizr.promises
}

document.addEventListener('DOMContentLoaded', () => {
    if (!meetsRequirements()) {
        return
    }

    const main = document.querySelector('main')
    const outerSection = document.getElementById('outer-wrapper')
    const sand = document.getElementById('sand')

    const easel = Easel.create(outerSection)

    if (Balloons.meetsRequirements()) {
        initBalloons(easel)
    }

    const sun = Sun.create(easel)

    if (Router.meetsRequirements()) {
        initRouting(main)
    }

    const mountain = desert.createMountain()
    document.body.insertBefore(mountain, sand)

    const desertFunc = desert.style(outerSection, 300)

    document.addEventListener('click', e => {
        const button = util.findAncestor(e.target, element => element.classList.contains('js-fillNow'))
        if (button != null) {
            const id = button.dataset.target
            const input = document.getElementById(id)
            util.setDateTimeInputToNow(input)
        }
    })

    window.addEventListener('load', () => {
        desertFunc().then(show(sun, mountain, sand))
        window.addEventListener('resize', desertFunc)
    })
})

function show(...elms) {
    return () => {
        requestAnimationFrame(() => {
            for(let elm of elms) {
                elm.classList.add('shown')
            }
        })
    }
}

const balloons = {
    MAX_AMOUNT: 10,
    EVERY: 20000,
    MIN_HEIGHT: .075,
    AVG_HEIGHT: .1,
    MAX_HEIGHT: .15
}

function initBalloons(easel) {
    Bliss.fetch('/dist/images/balloon.svg').then(svg => {
        const intID = setInterval(addBalloon, balloons.EVERY)
        requestAnimationFrame(addBalloon)

        function addBalloon() {
            if (Bliss.$('.svg-balloon', easel).length >= balloons.MAX_AMOUNT || (document.hidden || document.msHidden || document.webkitHidden)) {
                return
            }

            const screen = util.screenSize()

            const max = screen.width * balloons.MAX_HEIGHT
            const avg = screen.width * balloons.AVG_HEIGHT
            const min = screen.width * balloons.MIN_HEIGHT
            const bottom = screen.height - max

            const top = util.choosePoint(gaussRandom(), 0, bottom / 2, bottom)
            const left = screen.width
            const hexColor = Please.make_color({
                saturation: .8 + Math.random() * .2,
                value: .8 + Math.random() * .2
            })[0]
            const height = util.choosePoint(gaussRandom(), min, avg, max) * adjustHeigth(top, bottom)
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

        function adjustHeigth(top, bottom) {
            return Math.sqrt(1 - Math.pow(top / bottom, 2))
        }
    })
}

function initRouting(main) {
    Router.init()

    const outFunc = transitionOut(main)
    const inFunc = transitionIn(main)

    const handleForm = Router.handleForm(outFunc, inFunc)
    const handleRouting = Router.handleRouting(outFunc, inFunc)
    const handleBack = Router.handleBack(main, () => Promise.resolve(), (state) => {
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
        main.appendChild(frag)

        let targets = Array.from(main.children).reduce((list, e) => list.concat(transitionTarget(e)), [])
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

        return a.finished.then(() => {
            window.dispatchEvent(new Event('resize'))
        })
    }
}