const Please = require('pleasejs')
const sdRand = require('gauss-random')
const anime = require('animejs')

const util = require("./util")

exports.prepare = (svg) => {
	return new Promise((resolve, reject) => {
		if(DOMParser == null) {
			reject("DOMParser unavailable")
			return
		}

		const MAX_HEIGHT = getBalloonHeight(MAX_HEIGHT_N)

		const parser = new DOMParser()

		const dom = parser.parseFromString(svg.responseText, "image/svg+xml")
		const root = dom.documentElement

		const aspectRatio = root.width.baseVal.value / root.height.baseVal.value

		const screen = util.screenSize()

		const height = getBalloonHeight(sdRand());
		const width = aspectRatio * height;
		const hHalf = screen.height / 2;

		console.log(width, height)

		const position = {
			top: util.inBetween(hHalf + hHalf * sdRand(), 150, screen.height - 150),
			left: screen.width + width
		}
		
		const speed = Math.sqrt(width*width*width) / 2;
		const color = Please.make_color({
			saturation: .8 + Math.random() * .2,
			value: .8 + Math.random() * .2
		})[0];

		Bliss.style(root, {
			width: width + "px",
			height: "auto",
			position: "fixed",
			top: position.top + "px",
			left: 0 + "px",
			transform: `translateX(${position.left}px) translateY(0)`
		});

		const ids = new Map()
		Bliss.$("[id]", dom)
			.map(element => {return { dom: element, id: element.id}})
			.forEach(e => {
				ids.set(e.id, e.dom)
				e.dom.classList.add(e.id)
				e.dom.removeAttribute("id")
			})

		ids.get("flame").style.opacity = 1;
		ids.get("balloon").style.fill = color;

		resolve({
			dom,
			width,
			speed,
			position,
			root,
			ids
		})
	})
}

exports.drawOn = function (easel) {
	return (balloon) => {
		return new Promise((resolve, reject) => {
			easel.appendChild(balloon.root)
			resolve(balloon)
		})
	}
}

exports.fly = function (balloon) {
	ascend(balloon)
}

function descend(balloon) {
	if(balloon.position.left < -balloon.width) {
		balloon.root.remove();
		return;
	}

	balloon.position.left -= balloon.speed

	anime({
		targets: balloon.root,
		translateX: balloon.position.left + 'px',
		translateY: '0',
		easing: 'easeOutQuad',
		duration: 20000,
		complete: () => ascend(balloon)
	});
}

function ascend(balloon) {
	if(balloon.position.left < -balloon.width) {
		balloon.root.remove();
		return;
	}

	const climb = (balloon.speed + 100 + 50 * util.inBetween(sdRand(), -1, 1)) / 20;
	balloon.position.left -= balloon.speed/5;

	anime({
		targets: balloon.root,
		translateX: balloon.position.left + 'px',
		translateY: `-${climb}px`,
		easing: 'easeInQuad',
		duration: 4000,
		complete: () => descend(balloon)
	});

	/*var flame = balloon.ids.get("flame");
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
	});*/
}

const MAX_HEIGHT_N = 2
function getBalloonHeight(seed) {
	return 100 + 40 * util.inBetween(seed, -MAX_HEIGHT_N, MAX_HEIGHT_N)
}