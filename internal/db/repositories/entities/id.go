/*
 * Copyright (C) 2023 by Enterprise Technology, Viet Thai International
 * All Rights Reserved.
 *
 * This source code is protected under international copyright law.  All rights
 * reserved and protected by the copyright holders.
 * This file is confidential and only available to authorized individuals with the
 * permission of the copyright holders.  If you encounter this file and do not have
 * permission, please contact the copyright holders and delete this file.
 */

package entities

import "github.com/phuchnd/simple-go-boilerplate/internal/generators"

var (
	defaultIDGenerator IDGenerator
)

func init() {
	defaultIDGenerator, _ = NewIDGenerator(nil)
}

type IDGenerator interface {
	Next() ID
}

type idGeneratorImpl struct {
	sf generators.ISnowflakeIDGenerator
}

func NewIDGenerator(sf generators.ISnowflakeIDGenerator) (IDGenerator, error) {
	if sf == nil {
		var err error
		if sf, err = generators.NewSnowflakeIDGenerator(); err != nil {
			return nil, err
		}
	}
	return &idGeneratorImpl{sf}, nil
}

// Next returns
func (impl *idGeneratorImpl) Next() ID {
	return ID(impl.sf.Next())
}

// NextID returns a new ID using default generator.
func NextID() ID {
	return defaultIDGenerator.Next()
}
