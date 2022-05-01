module.exports = {
    future: {
        purgeLayersByDefault: true,
    },
    purge: {
        enabled: process.env.NODE_ENV === 'production',
        content: [
            'Shared/**/*.vue',
            'Pages/**/*.vue'
        ]
    },
    theme: {
        extend: {
            opacity: {
                '15': '0.15',
                '30': '0.30',
                '40': '0.40',
                '65': '0.65',
                '80': '0.80',
            }
        },
    },
    variants: {},
    plugins: [],
}