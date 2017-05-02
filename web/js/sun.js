module.exports = {
    create,
    setColor
}

function create(size) {
    const radius = size / 2

    const sun = document.createElement("div")
    sun.classList.add("sun")

    sun.style.position = 'fixed'
    sun.style.top = 0
    sun.style.left = 0

    sun.style.margin = `-${radius}px 0 0 -${radius}px`
    sun.style.height = `${size}px`
    sun.style.width = `${size}px`

    return sun
}

function setColor(sun, sunColor, raysColor) {
    sun.style.background = `radial-gradient(circle closest-side at center, ${sunColor} 0%, ${raysColor} 10%, transparent 100%)`

    return sun
}