import setuptools

with open("README.md", "r") as fh:
    long_description = fh.read()

setuptools.setup(
    name="ida",
    version="0.3.0",
    author="Aditya Duri",
    author_email="aditya.duri@gmail.com",
    description="Process voters for music competitions.",
    long_description=long_description,
    long_description_content_type="text/markdown",
    url="https://github.com/Iune/ida",
    packages=setuptools.find_packages(),
    classifiers=[
        "Programming Language :: Python :: 3",
        "License :: OSI Approved :: BSD License",
        "Operating System :: OS Independent",
    ],
    python_requires='>=3.6',
    install_requires=["pyperclip", "colorama", "termcolor"],
    entry_points={
        "console_scripts": [
            "ida=ida.ida:main"
        ]
    }
)