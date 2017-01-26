require("./shim")

const Anime = require("animejs")
const Coordinates = require('./coordinate')
const Balloons =  require('./balloons')
const Stick = require('./stick')
const util = require("./util")
const router = require("./router")
const Sun = require("./sun")

exports.debug = {
	Anime
}

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

	initEasel(easel)

	initBalloons(easel)

	initSun(easel)

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

		if(a == null || !router.meetRequirements()) {
			return;
		}

		e.preventDefault()

		const cleanUp = new Promise((resolve, reject) => {
			Anime({
				targets: main.children,
				duration: 500,
				easing: "linear",
				translateY: (_, i) => `-${(i+1) * 100}px`,
				opacity: 0,
				complete: () => {
					Array.from(main.children).forEach(child => main.removeChild(child))
					window.dispatchEvent(new Event("resize"))
					resolve()
				}
			})
		})

		Promise.all([
			cleanUp,
			router.next(a)
		]).then(([_, frag]) => {
			Array.from(frag.children).forEach((c, i) => {
				c.style.transform = `translateY(-${(i+1) * 100}px)`
				c.style.opacity = "0"
			})
			Anime({
				targets: frag.children,
				duration: 500,
				easing: "linear",
				translateY: "0",
				opacity: 1,
			})
			main.appendChild(frag)
		})
	})
	window.addEventListener("popstate", e => {
		main.innerHTML = e.state
		window.dispatchEvent(new Event("resize"))
	})
})

function initEasel(easel) {
	easel.style.position = 'fixed'
	easel.style.top = 0
}

function initSun(easel) {
	const sun = Sun.create(1000)
	Sun.setColor(sun, "#FFF250", "#FFFFFF")
	easel.appendChild(sun)
}


const balloons = {
	MAX: 20,
	EVERY: 5000
}
function initBalloons(easel) {
	const getBalloon = Bliss.fetch("/dist/svg/balloon.svg")
	setInterval(() => {
		if(easel.childNodes.length >= balloons.MAX || (document.hidden || document.msHidden || document.webkitHidden)) {
			return
		}
		getBalloon.then(Balloons.prepare)
			.then(Balloons.drawOn(easel))
			.then(Balloons.fly)
	}, balloons.EVERY)
}