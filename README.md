# statusupdate

[![Documentation](https://godoc.org/github.com/induzo/statusupdate?status.svg)](http://godoc.org/github.com/induzo/statusupdate) [![Go Report Card](https://goreportcard.com/badge/github.com/induzo/statusupdate)](https://goreportcard.com/report/github.com/induzo/statusupdate) [![Maintainability](https://api.codeclimate.com/v1/badges/3e4ee9ac6a7a39a18c36/maintainability)](https://codeclimate.com/github/induzo/statusupdate/maintainability) [![Coverage Status](https://coveralls.io/repos/github/induzo/statusupdate/badge.svg?branch=master)](https://coveralls.io/github/induzo/statusupdate?branch=master) [![CircleCI](https://circleci.com/gh/induzo/statusupdate.svg?style=svg)](https://circleci.com/gh/induzo/statusupdate)

rest way of dealing with status updates (through PATCH method)

## URIs

the URI should be of the form /entity/{ID}/statusupdate/{ActionID}

Using fsm, it will update the entity to the new status.
