require("./shim")

const Anime = require("animejs")
const Coordinates = require('./coordinate')
const Balloons =  require('./balloons')
const Stick = require('./stick')
const util = require("./util")
const router = require("./router")

exports.imports = {
	Coordinates,
	Anime
}

const balloons = exports.balloons = {}
balloons.MAX = 20
balloons.EVERY = 5000

function styleDesert(mainCloud) {
	mainCloud.style.marginBottom = 0
	return new Promise((resolve) => {
		setTimeout(() => {
			let screen = util.screenSize()
			mainCloud.style.marginBottom = (screen.height - 300) + "px"
			resolve()
		}, 0)
	})
}

document.addEventListener('DOMContentLoaded', () => {
	router.init()

	const main = Bliss("main")
	const easel = Bliss("#sky-easel")

	easel.style.position = 'fixed'
	easel.style.top = 0
	const getBalloon = Bliss.fetch("/dist/svg/balloon.svg")
	setInterval(() => {
		if(easel.childNodes.length >= balloons.MAX || (document.hidden || document.msHidden || document.webkitHidden)) {
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

	let mainCloud = Bliss("body > .cloud")
	let styleFunc = styleDesert.bind(null, mainCloud)

	let mountainRange = document.getElementById("mountain-range");
	let stickFunc = Stick.toBottom(mountainRange);
	let raiseEasel = () => {
		var height = util.screenSize().height
		const offset = Bliss(".sand").getBoundingClientRect().top
		if(offset < height) {
			easel.style.transform = `translateY(-${height - offset}px)`
		} else {
			easel.style.transform = null
		}
	}

	styleFunc()
		.then(stickFunc)
		.then(raiseEasel)
	window.addEventListener('load', () => {
		styleFunc()
			.then(stickFunc)
			.then(raiseEasel)
	})
	window.addEventListener('resize', () => {
		styleFunc()
			.then(stickFunc)
			.then(raiseEasel)
	})
	document.addEventListener('scroll', () => {
		stickFunc().then(raiseEasel)
	})
	window.addEventListener('click', e => {
		if(e.defaultPrevented) {
			return;
		}

		let a = router.firstA(e.target)
		if(a != null && router.meetRequirements()) {
			e.preventDefault()
			const cleanUp = new Promise((resolve, reject) => {
				Array.from(main.children).forEach(child => main.removeChild(child))
				resolve()
			})
			Promise.all([
				cleanUp,
				router.next(a)
			]).then(([_, content]) => {
				main.innerHTML = content
			})
		}
	})
	window.addEventListener("popstate", e => main.innerHTML = e.state)
})

const pushState = window.history.pushState
window.history.pushState = (...args) => {
	pushState.apply(window.history, args)
}