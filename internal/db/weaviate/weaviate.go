package db

import (
    "context"
    "fmt"
    "time"
)

func (w *WeaviateClient) AddMemory(ctx context.Context, content string, userID string) error {
    // FORCE a value if userID is empty to ensure summarization works
    if userID == "" {
        userID = "emmanuel"
    }

    properties := map[string]interface{}{
        "content":    content,
        "userId":     userID,   
        "timestamp":  time.Now().Format(time.RFC3339),
        "memoryType": "raw",
        "isSummary":  false,
    }

    // This log will appear in your terminal to confirm the save
    fmt.Printf("--- [DB] SAVING MEMORY FOR USER: %s ---\n", userID)

    _, err := w.client.Data().Creator().
        WithClassName("Memory_idx").
        WithProperties(properties).
        Do(ctx)
    
    if err != nil {
        return fmt.Errorf("failed to create memory: %w", err)
    }
    
    return nil
}