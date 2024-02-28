/*
Copyright (c) 2024 Sebastian Kroczek <me@xbug.de>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package simulation

import "github.com/gotameme/core/ant"

type SimulationConfig struct {
	// antDesiredValue defines how many ants should be in the simulation simultaneously
	antDesiredValue int
	// sugarDesiredValue defines how many sugar should be in the simulation simultaneously
	sugarDesiredValue int
	// antConstructor is a function that creates a new ant
	antConstructor ant.AntConstructor
	// chooseRole is a function that determines the role of an ant
	chooseRole ant.ChooseRole
	// roles is a map of roles and their properties
	roles                 map[string]Properties
	defaultRoleProperties Properties
}

func NewSimulationConfig(options ...SimulationOptions) *SimulationConfig {
	s := &SimulationConfig{
		antDesiredValue:   100,
		sugarDesiredValue: 1,
		antConstructor: func(os ant.AntOs) interface{} {
			return &struct{}{}
		},
		chooseRole: func(rolesCount ant.RolesCount) string {
			return ""
		},
		// TODO: We need the real default range, which is half the diagonal of the field
		defaultRoleProperties: NewDefaultProperties(100),
	}

	for _, o := range options {
		o(s)
	}

	return s
}

type SimulationOptions func(*SimulationConfig)

func WithAntDesiredValue(desiredValue int) SimulationOptions {
	return func(s *SimulationConfig) {
		s.antDesiredValue = desiredValue
	}
}

func WithSugarDesiredValue(desiredValue int) SimulationOptions {
	return func(s *SimulationConfig) {
		s.sugarDesiredValue = desiredValue
	}
}

func WithAntConstructor(antConstructor ant.AntConstructor) SimulationOptions {
	return func(s *SimulationConfig) {
		s.antConstructor = antConstructor
	}
}

func WithRoles(roles map[string]Properties, chooseRole ant.ChooseRole) SimulationOptions {
	return func(s *SimulationConfig) {
		s.roles = roles
		s.chooseRole = chooseRole
	}
}
