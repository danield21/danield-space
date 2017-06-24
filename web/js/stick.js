const util = require('./util')

exports.toBottom = (element, _up = 0) => {
    const wrapper = document.createElement('div')
    element.parentNode.insertBefore(wrapper, element)
    let frame

    return () => {
        if (frame) {
            cancelAnimationFrame(frame)
        }
        const rect = wrapper.getBoundingClientRect()
        const elementRect = element.getBoundingClientRect()
        const size = util.screenSize(true)

        if (rect.top + elementRect.height > size.height) {
            frame = requestAnimationFrame(() => {
                frame = null
                element.style.position = 'fixed'
                element.style.bottom = 0
                wrapper.style.height = elementRect.height + 'px'
            })
        } else {
            frame = requestAnimationFrame(() => {
                frame = null
                element.style.position = 'relative'
                element.style.bottom = null
                wrapper.style.height = 0
            })
        }
        return Promise.resolve()
    }
}