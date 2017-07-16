const anime = require('animejs')

const util = require('./util')
const stick = require('./stick')

function style(mainCloud, minAltitude) {
    return () => {
        return new Promise((resolve) => {
            let screen = util.screenSize()
            let cloudHeight = mainCloud.getBoundingClientRect().height
            let altitude = screen.height - cloudHeight
            if (altitude < minAltitude) {
                altitude = minAltitude
            }
            mainCloud.style.marginBottom = altitude + 'px'
            resolve()
        })
    }
}

function display(mountain, desert, sun) {
    return () => {
        mountain.classList.add('shown')
        desert.classList.add('shown')
        sun.classList.add('shown')
    }
}

function stickMountain(mountain) {
    return stick.toBottom(mountain)
}

module.exports = {
    style,
    display,
    stickMountain
}