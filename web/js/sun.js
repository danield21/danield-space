export function create(easel) {
    const sun = document.createElement('div')
    sun.setAttribute('id', 'sun')
    easel.append(sun)

    return sun
}
