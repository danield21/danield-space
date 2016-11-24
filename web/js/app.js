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
	window.addEventListener('WebComponentsReady', () => {
		let mainCloud = Bliss("body > nm-cloud")
		let styleFunc = styleDesert.bind(null, mainCloud)

		let mountainRange = document.getElementById("mountain-range");
		let stickFunc = Stick.toBottom(mountainRange);

		styleFunc().then(stickFunc)
		
		window.addEventListener('resize', () => {
			styleFunc().then(stickFunc)
		})
		document.addEventListener('scroll', () => { stickFunc() })

	})

	const path = Bliss("#js-balloon-flight")
	setInterval(() => {
		if(path.children.length < balloons.MAX) {
			Balloons.createFlying(path)
		}
	}, balloons.EVERY)
})