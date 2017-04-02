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

function choosePoint(z, min, mean, max) {
	var point;
	if(z > 0) {
		var dif = max - mean
		point = mean + (dif * (z/(z+1)))
	} else if(z < 0) {
		var dif = mean - min
		var pZ = -z
		point = min + (dif * (pZ/(pZ+1)))
	} else {
		point = mean
	}
	console.log("z: %s, min: %s, mean: %s, max: %s, point: %s", z, min, mean, max, point)
	return point;
}

function setDateTimeInputToNow(input) {
	var now = new Date();
	input.value = new Date(now.getTime()-now.getTimezoneOffset()*60000).toISOString().substring(0,19)
}

module.exports = {
	inBetween,
	screenSize,
	setDateTimeInputToNow,
	choosePoint
}