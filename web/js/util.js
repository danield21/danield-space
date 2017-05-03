function inBetween(value, min, max) {
    return Math.min(Math.max(value, min), max)
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

module.exports = {
    inBetween,
    screenSize,
    setDateTimeInputToNow,
    choosePoint
}