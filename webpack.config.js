var webpack = require('webpack'),
	path = require('path'),
	WebpackCopy = require('copy-webpack-plugin')
	ExtractTextPlugin = require("extract-text-webpack-plugin");

module.exports = {
	debug: true,
	entry: {
		app: './web/bundler.js',
		styles: './web/sass/app.scss'
	},
	output: {
		path: path.join(__dirname, 'app', 'dist'),
		filename: 'js/[name].js',
		library: '[name]'
	},
	module: {
		loaders: [
			{
				test: /\.js$/,
				loader: "babel-loader",
				query: {
					presets: ['es2015']
				}
			},
			{
				test: /\.scss$/,
				loader: ExtractTextPlugin.extract("style-loader", "css-loader!sass-loader")
			}
		]
	},
	plugins: [
		new WebpackCopy([
			{from: path.join("web", "view"), to: path.join("..", "view")},
			{from: path.join("web", "svg"), to: path.join("svg")},
			{from: path.join("web", "lib"), to: path.join("js")},
			{from: path.join("web", "bower_components", "webcomponentsjs", "webcomponents-lite.min.js"), to: path.join("js", "webcomponents-lite.min.js")}
		]),
		new ExtractTextPlugin("css/[name].css")
	]
};