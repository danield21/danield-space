import * as util from './util'

export function style(mainCloud, minAltitude) {
    return () => {
        return new Promise((resolve) => {
            let screen = util.screenSize()
            let cloudHeight = mainCloud.getBoundingClientRect().height
            let altitude = screen.height - cloudHeight
            if (altitude < minAltitude) {
                altitude = minAltitude
            }
            mainCloud.lastElementChild.style.marginBottom = altitude + 'px'
            resolve()
        })
    }
}

export function createMountain() {
    const embed = document.createElement('embed')
    embed.setAttribute('id', 'mountain-range')
    embed.setAttribute('src', '/dist/images/mountain.svg')
    return embed
}
