function inBetween(value, min, max) {
	return Math.min(Math.max(value, min), max);
}

function screenSize() {
	return {
		width: window.innerWidth
			|| document.documentElement.clientWidth
			|| document.body.clientWidth,
		height: window.innerHeight
			|| document.documentElement.clientHeight
			|| document.body.clientHeight
	}
}

module.exports = {
	inBetween,
	screenSize
}