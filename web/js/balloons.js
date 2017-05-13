const Please = require('pleasejs')
const sdRand = require('gauss-random')
const anime = require('animejs')
const Modernizr = require('modernizr')

const util = require('./util')

module.exports = {
    prepare,
    drawOn,
    fly,
    meetsRequirements
}

function meetsRequirements() {
    return Modernizr.inlinesvg &&
        Modernizr.requestanimationframe &&
        Modernizr.classlist &&
        window.DOMParser
}

const MIN_HEIGHT = 100
const MAX_HEIGHT = 300
const AVG_HEIGHT = 200

function prepare(svg) {
    return new Promise((resolve, _reject) => {

        const parser = new DOMParser()

        const dom = parser.parseFromString(svg.responseText, 'image/svg+xml')
        const root = dom.documentElement

        const aspectRatio = root.width.baseVal.value / root.height.baseVal.value

        const screen = util.screenSize()

        const height = getBalloonHeight(sdRand())
        const width = aspectRatio * height
        const hHalf = screen.height / 2

        const position = {
            top: util.choosePoint(sdRand(), 0, hHalf, screen.height, MAX_HEIGHT),
            left: screen.width + width
        }

        const speed = Math.sqrt(width * width * width) / 4
        const color = Please.make_color({
            saturation: .8 + Math.random() * .2,
            value: .8 + Math.random() * .2
        })[0]

        Bliss.style(root, {
            width: width + 'px',
            height: 'auto',
            position: 'fixed',
            top: position.top + 'px',
            left: 0 + 'px',
            transform: `translateX(${position.left}px) translateY(0)`
        })

        const ids = new Map()
        Bliss.$('[id]', dom)
            .map(element => { return { dom: element, id: element.id } })
            .forEach(e => {
                ids.set(e.id, e.dom)
                e.dom.classList.add(e.id)
                e.dom.removeAttribute('id')
            })

        ids.get('flame').style.opacity = 1
        ids.get('balloon').style.fill = color

        resolve({
            dom,
            width,
            speed,
            position,
            root,
            ids
        })
    })
}

function drawOn(easel) {
    return (balloon) => {
        return new Promise((resolve, _reject) => {
            easel.appendChild(balloon.root)
            resolve(balloon)
        })
    }
}

function fly(balloon) {
    ascend(balloon)
}

function descend(balloon) {
    if (balloon.position.left < -balloon.width) {
        balloon.root.remove()
        return
    }

    balloon.position.left -= balloon.speed

    anime({
        targets: balloon.root,
        translateX: balloon.position.left + 'px',
        translateY: '0',
        easing: 'easeOutQuad',
        duration: 20000,
        complete: () => ascend(balloon)
    })
}

function ascend(balloon) {
    if (balloon.position.left < -balloon.width) {
        balloon.root.remove()
        return
    }

    const climb = (balloon.speed + 100 + 50 * util.inBetween(sdRand(), -1, 1)) / 20
    balloon.position.left -= balloon.speed / 5

    anime({
        targets: balloon.root,
        translateX: balloon.position.left + 'px',
        translateY: `-${climb}px`,
        easing: 'easeInQuad',
        duration: 4000,
        complete: () => descend(balloon)
    })
}

function getBalloonHeight(seed) {
    return util.choosePoint(seed, MIN_HEIGHT, AVG_HEIGHT, MAX_HEIGHT)
}