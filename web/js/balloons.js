const anime = require('animejs')
const Modernizr = require('modernizr')

module.exports = {
    meetsRequirements,
    parseSVG,
    size,
    position,
    storeIds,
    color,
    drawOn,
    fly,
    remove
}

function meetsRequirements() {
    return Modernizr.inlinesvg &&
        Modernizr.requestanimationframe &&
        Modernizr.classlist &&
        window.DOMParser
}

function parseSVG(svg) {
    return new Promise((resolve, _reject) => {
        const parser = new DOMParser()

        const dom = parser.parseFromString(svg.responseText, 'image/svg+xml')
        const root = dom.documentElement
        resolve({ dom, root })
    })
}

function size(height) {
    return function(balloon) {
        let aspectRatio = balloon.root.width.baseVal.value / balloon.root.height.baseVal.value
        let width = height * aspectRatio

        Bliss.style(balloon.root, {
            width: width + 'px',
            height: 'auto'
        })

        Object.defineProperties(balloon, {
            height: { value: height },
            width: { value: width },
            aspectRatio: { value: aspectRatio },
            speed: { value: Math.sqrt(width * width * width) / 4 }
        })

        return Promise.resolve(balloon)
    }
}

function position(top, left) {
    return function(balloon) {
        left += balloon.width

        Bliss.style(balloon.root, {
            position: 'fixed',
            top: top + 'px',
            left: 0 + 'px',
            transform: `translateX(${left}px) translateY(0)`
        })

        balloon.position = { top, left }

        return Promise.resolve(balloon)
    }
}

function storeIds(balloon) {
    const ids = new Map()
    Bliss.$('[id]', balloon.dom)
        .map(element => { return { dom: element, id: element.id } })
        .forEach(e => {
            ids.set(e.id, e.dom)
            e.dom.classList.add(e.id)
            e.dom.removeAttribute('id')
        })
    balloon.ids = ids
    return Promise.resolve(balloon)
}

function color(hexColor) {
    return function(balloon) {
        balloon.ids.get('balloon').style.fill = hexColor
        return Promise.resolve(balloon)
    }
}

function drawOn(easel) {
    return (balloon) => {
        easel.appendChild(balloon.root)
        return Promise.resolve(balloon)
    }
}

function remove(easel) {
    return (balloon) => {
        easel.removeChild(balloon.root)
        return Promise.resolve(balloon)
    }
}

function fly(balloon) {
    return ascend(balloon).apply()
}

function descend(balloon) {
    return function() {
        if (balloon.position.left < -balloon.width) {
            return Promise.resolve(balloon)
        }

        balloon.position.left -= balloon.speed

        const a = anime({
            targets: balloon.root,
            translateX: balloon.position.left + 'px',
            translateY: '0',
            easing: 'easeOutQuad',
            duration: 20000
        })

        return a.finished.then(ascend(balloon))
    }
}

function ascend(balloon) {
    return function() {
        if (balloon.position.left < -balloon.width) {
            return Promise.resolve(balloon)
        }

        const climb = (balloon.speed + 100) / 20
        balloon.position.left -= balloon.speed / 5

        const a = anime({
            targets: balloon.root,
            translateX: balloon.position.left + 'px',
            translateY: `-${climb}px`,
            easing: 'easeInQuad',
            duration: 4000,
            complete: () => descend(balloon)
        })

        return a.finished.then(descend(balloon))
    }
}