package main

// struct of the config model
type Configuration struct {
	Proxy ProxyConfig
}

type ProxyConfig struct {
	Listen   ListenConfig
	Services []ServicesConfig
}

type ListenConfig struct {
	Port    int
	Address string
}

type ServicesConfig struct {
	Name   string
	Domain string
	Hosts  []ServiceHostConfig
}

type ServiceHostConfig struct {
	Address string
	Port    int
}
