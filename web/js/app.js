const imports = exports.imports = {}

const Coordinates = imports.Coordinates = require('./coordinate')
const Balloons = imports.Balloons = require('./balloons')
const util = require("./util")

const balloons = exports.balloons = {}
balloons.MAX = 20
balloons.EVERY = 5000

function styleDesert(mainCloud) {
	mainCloud.style.marginBottom = 0
	setTimeout(() => {
		let screen = util.screenSize()
		let dom = document.body.getBoundingClientRect()
		let offset = ((dom.height+300 < screen.height) ? (screen.height - dom.height) : 0) + 300
		mainCloud.style.marginBottom = offset + "px"
		console.log(offset)
	}, 0)
}   

document.addEventListener('DOMContentLoaded', () => {
	window.addEventListener('WebComponentsReady', () => {
		let mainCloud = Bliss("body > nm-cloud")
		let styleFunc = styleDesert.bind(null, mainCloud)
		styleFunc()
		window.addEventListener('resize', styleFunc)
	})

	const path = Bliss("#js-balloon-flight")
	setInterval(() => {
		if(path.children.length < balloons.MAX) {
			Balloons.createFlying(path)
		}
	}, balloons.EVERY)
})