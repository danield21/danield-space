export function create() {
    const easel = document.createElement('div')
    easel.setAttribute('id', 'sky-easel')
    document.body.insertBefore(easel, document.body.firstChild)

    return easel
}