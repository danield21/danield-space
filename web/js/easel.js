export function create(target) {
    const canvas = document.createElement('div')
    canvas.setAttribute('id', 'sky-canvas')
    target.parentElement.insertBefore(canvas, target.nextSibling)

    const easel = document.createElement('div')
    easel.setAttribute('id', 'sky-easel')
    canvas.append(easel)

    return easel
}