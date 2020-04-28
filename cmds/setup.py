# coding: utf-8
import os
import sys

from setuptools import find_packages, setup

def check_requires():
    # parse file location in, check requiremnet file.
    requirements = []
    for line in open('test1.txt').readlines():
        if line.startswith('#') or line == '' or line.startswith('http') or line.startswith('.'):
            continue
        requirements.append(line)
    return requirements

setup(
    name='local_topo',
    version=get_version(),
    license='MIT',
    author='Danni Li',
    author_email='dannili3-c@my.cityu.edu.hk',
    description='Local Cluster Set up via user-defined topology graph',
    packages=find_packages(exclude=['tests']),
    install_requires=get_install_requires(),
    include_package_data=True
)

