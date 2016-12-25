require("./shim")

const imports = exports.imports = {}

const Coordinates = imports.Coordinates = require('./coordinate')
const Balloons = imports.Balloons = require('./balloons')
const Stick = imports.Stick = require('./stick')
const util = require("./util")

const balloons = exports.balloons = {}
balloons.MAX = 20
balloons.EVERY = 5000

function styleDesert(mainCloud) {
	mainCloud.style.marginBottom = 0
	return new Promise((resolve) => {
		setTimeout(() => {
			let screen = util.screenSize()
			let dom = document.body.getBoundingClientRect()
			let offset = ((dom.height+300 < screen.height) ? (screen.height - dom.height) : 0) + 300
			mainCloud.style.marginBottom = offset + "px"
			resolve()
		}, 0)
	})
}

document.addEventListener('DOMContentLoaded', () => {

	const easel = Bliss("#sky-easel")
	easel.style.position = 'fixed'
	easel.style.top = 0
	const getBalloon = Bliss.fetch("dist/svg/balloon.svg")
	setInterval(() => {
		if(easel.childNodes.length >= balloons.MAX) {
			return
		}
		getBalloon.then(Balloons.prepare)
			.then(Balloons.drawOn(easel))
			.then(Balloons.fly)
	}, balloons.EVERY)

	;(() => {
		let sun = document.createElement("div")
		const size = 1000;
		const radius = size/2;
		sun.style.position = 'fixed'
		sun.style.top = 0
		sun.style.left = 0
		sun.style.margin = `-${radius}px 0 0 -${radius}px`
		sun.style.background = 'radial-gradient(circle closest-side at center, #FFFF00 0%, #FFFFFF 10%, transparent 100%)'
		sun.style.height = `${size}px`
		sun.style.width = `${size}px`
		easel.appendChild(sun);
	})();

	window.addEventListener('WebComponentsReady', () => {
		let mainCloud = Bliss("body > nm-cloud")
		let styleFunc = styleDesert.bind(null, mainCloud)

		let mountainRange = document.getElementById("mountain-range");
		let stickFunc = Stick.toBottom(mountainRange);

		styleFunc().then(stickFunc)
		
		window.addEventListener('resize', () => {
			styleFunc().then(stickFunc)
		})
		document.addEventListener('scroll', () => {
			stickFunc().then(() => {
				var height = util.screenSize().height
				const offset = Bliss(".sand").getBoundingClientRect().top
				if(offset < height) {
					easel.style.transform = `translateY(-${height - offset}px)`
				} else {
					easel.style.transform = null
				}
			})
		})

	})
})