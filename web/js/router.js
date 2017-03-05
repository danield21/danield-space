const Modernizr = require("modernizr")

module.exports = {
	firstA,
	init,
	next,
	meetsRequirements,
	handleRouting,
	handleForm,
}

function meetsRequirements() {
	return Modernizr.documentfragment &&
	Modernizr.xhrresponsetypedocument &&
	Modernizr.history
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

function firstForm(element) {
	if(element.tagName.toUpperCase() == "FORM") {
		return element
	}
	if(element.parentElement == null) {
		return null
	}
	return firstForm(element.parentElement)
}

function next(url, perform) {
	return perform.then(html => {
		window.history.pushState(html.body.innerHTML, window.document.title, url)
		const frag = document.createDocumentFragment()
		Array.from(html.body.children).forEach(c => frag.appendChild(c))
		return Promise.resolve(frag)
	}, e => {
		return Promise.reject(e)
	})
}

function submitForm(form) {
	var url = form.action
	var method = form.method
	var data = Array.from(form.elements).reduce((encode, e) => {
		if(e.name) {
			if(encode) encode += "&"
			return encode + encodeURIComponent(e.name) + "=" + encodeURIComponent(e.value)
		}
		return encode
	}, "")
	if(data !== "") {
		data += "&"
	}
	data += "theme=none"
	return Bliss.fetch(form.action, { responseType: "document", method, data}).then(response => {
		return response.responseXML ? Promise.resolve(response.responseXML) : Promise.reject(new Error("Did not get a document back from " + url + "?" + data))
	}, e => {
		return e.xhr.responseXML ? Promise.resolve(e.xhr.responseXML) : Promise.reject(new Error("Did not get a document back from " + url + "?" + data))
	})
}

function navigate(a) {
	if(a && a.tagName.toUpperCase() != "A" && a.href) {
		return Promise.reject(new Error("Provided value is not an A element with an href"))
	}
	var url = a.href

	var queryRegex = /\?[\w\d%=\[\]&]+$/
	if(url.match(queryRegex)) {
		url += "&"
	} else {
		url += "?"
	}
	url += "theme=none"

	return Bliss.fetch(url, { responseType: "document"}).then(response => {
		return response.responseXML ? Promise.resolve(response.responseXML) : Promise.reject(new Error("Did not get a document back from " + url))
	}, e => {
		return e.xhr.responseXML ? Promise.resolve(e.xhr.responseXML) : Promise.reject(new Error("Did not get a document back from " +  url))
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
			next(a.href, navigate(a))
		]).then(transitionIn)
	}
}

function handleForm(transitionOut, transitionIn) {
	return e => {
		if(e.defaultPrevented) {
			return;
		}

		let form = firstForm(e.target)

		if(form == null) {
			return;
		}

		e.preventDefault()

		Promise.all([
			transitionOut(),
			next(form.action, submitForm(form))
		]).then(transitionIn)
	}
}