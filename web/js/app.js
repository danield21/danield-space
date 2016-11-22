const imports = exports.imports = {}

const Coordinates = imports.Coordinates = require('./coordinate')
const Please = imports.Please = require('pleasejs')
const sdRand = imports.sdRand = require('gauss-random')
const anime = imports.anime = require('animejs')

exports.createBackgroundBalloon = (target) => {
	var balloon = document.createElement("nm-balloon");

	const screen = {
		width: window.innerWidth
			|| document.documentElement.clientWidth
			|| document.body.clientWidth,
		height: window.innerHeight
			|| document.documentElement.clientHeight
			|| document.body.clientHeight
	}

	const width = 100 + 50 * inBetween(sdRand(), -1, 1);
	const hHalf = screen.height / 2;

	const position = {
		top: hHalf + hHalf * inBetween(sdRand(), -1, 1),
		left: screen.width + width + (50 + 50 * inBetween(sdRand(), -1, 1))
	}
	
	const speed = Math.sqrt(width*width*width) / 2;
	const color = Please.make_color({
		saturation: .8 + Math.random() * .2,
		value: .8 + Math.random() * .2
	})[0];

	Bliss.style(balloon, {
		position: "fixed",
		top: position.top + "px",
		left: 0 + "px",
		transform: `translateX(${position.left}px) translateY(0)`
	});
	balloon.getSVGDocument().then(svg => {
		var flame = svg.getElementById("flame");
		flame.style.opacity = 0;
	})

	balloon.width = width;
	balloon.color = color;

	target.appendChild(balloon)
	
	setTimeout(ascendBalloon.bind(null, {
		dom: balloon,
		width,
		speed,
		position
	}), 10);
}

function inBetween(value, min, max) {
	return Math.min(Math.max(value, min), max);
}

function descendBalloon(balloon) {
	if(balloon.position.left < -balloon.width) {
		balloon.dom.remove();
		return;
	}

	balloon.position.left -= balloon.speed

	anime({
		targets: balloon.dom,
		translateX: balloon.position.left + 'px',
		translateY: '0',
		easing: 'easeOutQuad',
		duration: 20000,
		complete: () => ascendBalloon(balloon)
	});
}

function ascendBalloon(balloon) {
	if(balloon.position.left < -balloon.width) {
		balloon.dom.remove();
		return;
	}

	const climb = (balloon.speed + 100 + 50 * inBetween(sdRand(), -1, 1)) / 20;
	balloon.position.left -= balloon.speed/5;

	anime({
		targets: balloon.dom,
		translateX: balloon.position.left + 'px',
		translateY: `-${climb}px`,
		easing: 'easeInQuad',
		duration: 4000,
		complete: () => descendBalloon(balloon)
	});

	balloon.dom.getSVGDocument().then(svg => {
		var flame = svg.getElementById("flame");
		anime({
			targets: flame,
			opacity: 1,
			easing: 'easeInQuad',
			duration: 4000,
			complete: () => {
				anime({
					targets: flame,
					opacity: 0,
					easing: 'easeInQuad',
					duration: 2000
				});		
			}
		});
	})
}

setInterval(() => {
	var path = Bliss("#js-balloon-flight")
	const max = 20;
	if(path.children.length < max) {
		exports.createBackgroundBalloon(path)
	}
}, 5000)