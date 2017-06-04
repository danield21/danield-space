const util = require('./util')

exports.toBottom = (element, _up = 0) => {
    const wrapper = document.createElement('div')
    element.parentNode.insertBefore(wrapper, element.nextSibling)

    return () => {
        const rect = wrapper.getBoundingClientRect()
        const elementRect = element.getBoundingClientRect()
        const size = util.screenSize(true)

        if (rect.top + elementRect.height > size.height) {

            element.style.position = 'fixed'
            element.style.bottom = 0
            wrapper.style.height = elementRect.height + 'px'
        } else {
            element.style.position = 'relative'
            element.style.bottom = null
            wrapper.style.height = 0
        }
        return new Promise(resolve => { resolve() })
    }
}