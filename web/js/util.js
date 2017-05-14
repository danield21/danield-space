module.exports = {
    screenSize,
    setDateTimeInputToNow,
    choosePoint,
    scrollToTop,
    findAncestor
}

function screenSize() {
    return {
        width: window.innerWidth ||
            document.documentElement.clientWidth ||
            document.body.clientWidth,
        height: window.innerHeight ||
            document.documentElement.clientHeight ||
            document.body.clientHeight
    }
}

function choosePoint(z, min, mean, max) {
    if (z > 0) {
        let dif = max - mean
        return mean + (dif * (z / (z + 1)))
    } else if (z < 0) {
        let dif = mean - min
        let pZ = -z
        return min + (dif * (pZ / (pZ + 1)))
    } else {
        return mean
    }
}

function setDateTimeInputToNow(input) {
    var now = new Date()
    input.value = new Date(now.getTime() - now.getTimezoneOffset() * 60000).toISOString().substring(0, 19)
}

function scrollToTop(scrollDuration) {
    const scrollStep = -window.scrollY / (scrollDuration / 15)
    const scrollInterval = setInterval(function() {
        if (window.scrollY != 0) {
            window.scrollBy(0, scrollStep)
        } else clearInterval(scrollInterval)
    }, 15)
}

function findAncestor(element, match) {
    if (match(element)) {
        return element
    }
    if (element.parentElement == null) {
        return null
    }
    return findAncestor(element.parentElement)
}