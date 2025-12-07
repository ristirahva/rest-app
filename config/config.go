// config.go

package config

import (
    "encoding/json"
    "os"
)

type Config struct {
    Server struct {
        Host string `json:"host"`
        Port string `json:"port"`
    } `json:"server"`
    
    Database struct {
        Host     string `json:"host"`
        Port     string `json:"port"`
        Username string `json:"username"`
        Password string `json:"password"`
        Name     string `json:"name"`
    } `json:"database"`
    
    Logging struct {
        Level  string `json:"level"`
        Output string `json:"output"`
    } `json:"logging"`

    Validation struct{
        MinBarrelCapacity string `json:"min-barrel-capacity"`
        MaxBarrelCapacity string `json:"max-barrel-capacity"`
    } `json:"validation"`
}

// config.go (продолжение)
func LoadConfig(filename string) (*Config, error) {
    // Чтение файла
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    
    // Декодирование JSON
    var config Config
    if err := json.Unmarshal(data, &config); err != nil {
        return nil, err
    }
    
    return &config, nil
}
