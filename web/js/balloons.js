const Please = require('pleasejs')
const sdRand = require('gauss-random')
const anime = require('animejs')

const util = require("./util")

exports.createFlying = (target) => {
	var balloon = document.createElement("nm-balloon");

	const screen = util.screenSize()

	const width = 100 + 50 * util.inBetween(sdRand(), -1, 1);
	const hHalf = screen.height / 2;

	const position = {
		top: hHalf + hHalf * util.inBetween(sdRand(), -1, 1),
		left: screen.width + width + (50 + 50 * util.inBetween(sdRand(), -1, 1))
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
	
	setTimeout(ascend.bind(null, {
		dom: balloon,
		width,
		speed,
		position
	}), 10);
}

function descend(balloon) {
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
		complete: () => ascend(balloon)
	});
}

function ascend(balloon) {
	if(balloon.position.left < -balloon.width) {
		balloon.dom.remove();
		return;
	}

	const climb = (balloon.speed + 100 + 50 * util.inBetween(sdRand(), -1, 1)) / 20;
	balloon.position.left -= balloon.speed/5;

	anime({
		targets: balloon.dom,
		translateX: balloon.position.left + 'px',
		translateY: `-${climb}px`,
		easing: 'easeInQuad',
		duration: 4000,
		complete: () => descend(balloon)
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