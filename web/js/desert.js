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

function display(mountain, desert) {
    return () => {
        mountain.style.transform = 'translateY(100%)'
        mountain.style.visibility = 'visible'
        desert.style.transform = 'translateY(100%)'
        desert.style.visibility = 'visible'

        anime({
            targets: [mountain, desert],
            duration: 1000,
            easing: 'linear',
            translateY: 0,
        })
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