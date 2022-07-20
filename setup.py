import setuptools

setuptools.setup(
    version="0.0.1",
    license='mit',
    name="notify",
    author='nathan todd-stone',
    author_email='me@nathants.com',
    url='http://github.com/nathants/notify',
    install_requires=[
        'argh',
        'blessed',
    ],
    scripts=['notify'],
    description='notify',
)
