const Modernizr = require("modernizr")

module.exports = {
	firstA,
	getRoute,
	init,
	next,
	meetsRequirements,
	handleRouting
}

function meetsRequirements() {
	return Modernizr.history &&
	Modernizr.documentfragment &&
	Modernizr.xhrresponsetypedocument
}

function init() {
	const page = window.location.origin + window.location.pathname
	window.history.replaceState(Bliss("main").innerHTML, window.document.title, page)
}

function firstA(element) {
	if(element.tagName.toUpperCase() == "A") {
		return element
	}
	if(element.parentElement == null) {
		return null
	}
	return firstA(element.parentElement)
}

function next(a) {
	return getRoute(a).then(html => {
		window.history.pushState(html.body.innerHTML, window.document.title, a.href)
		const frag = document.createDocumentFragment()
		Array.from(html.body.children).forEach(c => frag.appendChild(c))
		return Promise.resolve(frag)
	}, e => {
		return Promise.reject(e)
	})
}

function getRoute(a) {
	if(a && a.tagName.toUpperCase() != "A" && a.href) {
		return Promise.reject(new Error("Provided value is not an A element with an href"))
	}

	return Bliss.fetch(a.href + "?theme=none", { responseType: "document"}).then(response => {
		return response.responseXML ? Promise.resolve(response.responseXML) : Promise.reject(new Error("Did not get a document back from " + a.href))
	}, e => {
		return e.xhr.responseXML ? Promise.resolve(e.xhr.responseXML) : Promise.reject(new Error("Did not get a document back from " + a.href))
	})
}

function handleRouting(transitionOut, transitionIn) {
	return e => {
		if(e.defaultPrevented) {
			return;
		}

		let a = firstA(e.target)

		if(a == null) {
			return;
		}

		e.preventDefault()

		Promise.all([
			transitionOut(),
			next(a)
		]).then(transitionIn)
	}
}