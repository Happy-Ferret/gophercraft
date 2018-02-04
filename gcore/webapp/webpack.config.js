var webpack = require('webpack');
const path = require("path");
module.exports = {
    entry: './main.tsx',
    output: {
        path: path.resolve(__dirname, './public'),  
        filename: 'main.js',
        publicPath: '/'
    },
    resolve: {
        extensions: ['.webpack.js', '.web.js', '.ts', '.tsx', '.js']
    },
    module: {
        loaders: [
            { test: /\.tsx?$/, loader: 'ts-loader' }
        ]
    },
    plugins: [
        new webpack.DefinePlugin({
            'process.env': {
                NODE_ENV: JSON.stringify('production')
            }
        }),
        new webpack.optimize.UglifyJsPlugin({ output: {comments: false} })
    ]
}