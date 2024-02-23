// Package datafile allows reading data samples from files
package main

import (
    "bufio"
    "os"
    "strconv"
    )

// GetFloats reads a float from each line of a file
func GetFloats(fileName string) ([]float64, error) {
    var numbers []float64
    file, err := os.Open(filename)
    if err != nil {
    return numbers, err
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
    numbers, err := stronv.ParseFloat(scanner.Text(), 64)
    if err != nil {
    return numbers, err
    }
    numbers = append(numbers, number)
    err := file.Close()
    if err != nil {
    return numbers, err
    }
    if scanner.Err() != nil {
        return numbers, scanner.Err()
        }
        return numbers, nil
    }