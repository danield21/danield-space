require("./shim")

const Anime = require("animejs")
const Coordinates = require('./coordinate')
const Balloons =  require('./balloons')
const Stick = require('./stick')
const util = require("./util")
const router = require("./router")
const Sun = require("./sun")
const Modernizr = require("modernizr")

exports.imports = {
	Anime,
	meetsRequirements,
	router
}

function meetsRequirements() {
	return Modernizr.eventlistener &&
	Modernizr.queryselector &&
	Modernizr.es5 &&
	Modernizr.promises
}

function styleDesert(mainCloud, desert) {
	mainCloud.style.marginBottom = 0
	desert.style.visibility = "hidden"
	return new Promise((resolve) => {
		setTimeout(() => {
			desert.style.visibility = "visible"
			let screen = util.screenSize()
			mainCloud.style.marginBottom = (screen.height - 300) + "px"
			resolve()
		}, 0)
	})
}

document.addEventListener('DOMContentLoaded', () => {
	if(!meetsRequirements()) {
		return
	}

	const main = Bliss("main")
	const easel = Bliss("#sky-easel")

	initEasel(easel)

	if(Balloons.meetsRequirements()) {
		initBalloons(easel)
	}

	initSun(easel)

	if(router.meetsRequirements()) {
		initRouting(main)
	}

	let mainCloud = Bliss("body > .cloud")
	let desert = Bliss("body > footer > .sand")
	let styleFunc = styleDesert.bind(null, mainCloud, desert)

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

	Bliss.$(".js-fillNow").forEach(button => {
		button.addEventListener('click', e => {
			var id = e.target.dataset.target
			var input = document.getElementById(id)
			util.setDateTimeInputToNow(input)
		})
	})

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
	MAX: 10,
	EVERY: 20000
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

function initRouting(main) {
	router.init()

	window.addEventListener('click', e => {
		const outStruct = transitionOut(main)
		const inStruct = transitionIn(main)
		router.handleRouting(outStruct.func, inStruct.func)(e)
	})

	window.addEventListener('submit', e => {
		const outStruct = transitionOut(main)
		const inStruct = transitionIn(main)
		router.handleForm(outStruct.func, inStruct.func)(e)
	})
	window.addEventListener("popstate", e => {
		main.innerHTML = e.state
		window.dispatchEvent(new Event("resize"))
	})
}

function transitionOut(main) {
	let transition
	const resolvable = {}
	const promise = new Promise((resolve, reject) => {
		resolvable.resolve = resolve;
		resolvable.reject = reject;
	})
	const func = () => {
		Anime({
			targets: main.children,
			duration: 500,
			easing: "linear",
			translateY: (_, i) => `-${(i+1) * 100}px`,
			opacity: 0,
			complete: () => {
				Array.from(main.children).forEach(child => main.removeChild(child))
				window.dispatchEvent(new Event("resize"))
				resolvable.resolve()
			}
		})
		return transition.promise
	}

	transition = { promise, func }

	return transition
}

function transitionIn(main) {
	let transition
	const resolvable = {}
	const promise = new Promise((resolve, reject) => {
		resolvable.resolve = resolve;
		resolvable.reject = reject;
	})
	const func = ([_, frag]) => {
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
			complete: () => {
				resolvable.resolve()
			}
		})
		main.appendChild(frag)
		return transition.promise
	}

	transition = { promise, func }

	return transition
}