var path = require('path');

module.exports = {
    mode: 'development',
    entry: './app/index.js',
    output: {
        filename: 'bundle.js',
        path: path.resolve(__dirname, 'dist')
    },
    module: {
        rules: [
          {
            test: /\.css$/,
            loader: ['style-loader', 'css-loader'], 
            include: [
                path.resolve(__dirname, 'node_modules'),
                path.resolve(__dirname, 'app/css'),
            ],
          },
        ],
    }
};