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

function setDateTimeInputToNow(input) {
	var now = new Date();
	input.value = new Date(now.getTime()-now.getTimezoneOffset()*60000).toISOString().substring(0,19)
}

module.exports = {
	inBetween,
	screenSize,
	setDateTimeInputToNow
}