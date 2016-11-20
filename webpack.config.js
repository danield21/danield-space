var webpack = require('webpack'),
	path = require('path'),
	WebpackCopy = require('copy-webpack-plugin');

module.exports = {
	debug: true,
	entry: {
		main: './web/bundler.js'
	},
	output: {
		path: path.join(__dirname, 'app', 'dist'),
		filename: 'bundled.js',
		library: 'app'
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
				loader: "style!css!sass"
			}
		]
	},
	plugins: [
		new WebpackCopy([
			{from: path.join("web", "view"), to: path.join("..", "view")},
			{from: path.join("web", "svg"), to: path.join("svg")}
		])
	]
};