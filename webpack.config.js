var webpack = require('webpack');
var path = require('path');
module.exports = {
    entry: [
      'webpack/hot/only-dev-server',
      "./assets/app.js"
    ],
    output: {
        path: path.join(__dirname, '/static'),
        filename: "bundle.js"
    },
    module: {
        loaders: [
            { test: /\.jsx?$/, loaders: ['babel'], exclude: /node_modules/ },
            { test: /\.js$/, exclude: /node_modules/, loader: 'babel-loader'},
            { test: /\.css$/, loader: "style!css" }
        ]
    },
    plugins: [
     new webpack.HotModuleReplacementPlugin(),
     new webpack.NoErrorsPlugin()
    ]

};
