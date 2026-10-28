				
				package main

				import (
					"crypto/sha256"
						"fmt"
							"time"
							)

							type Block struct {
								Index        int
									Timestamp    string
										Data         string
											PreviousHash string
												Hash         string
												}

												func calculateHash(b Block) string {
													record := fmt.Sprintf("%d%s%s%s",
															b.Index, b.Timestamp, b.Data, b.PreviousHash)
																sum := sha256.Sum256([]byte(record))
																	return fmt.Sprintf("%x", sum)
																	}

																	func createBlock(index int, data, previousHash string) Block {
																		b := Block{
																				Index:        index,
																						Timestamp:    time.Now().UTC().Format(time.RFC3339Nano),
																								Data:         data,
																										PreviousHash: previousHash,
																											}
																												b.Hash = calculateHash(b)
																													return b
																													}

																													func isChainValid(chain []Block) bool {
																														for i, b := range chain {
																																if b.Hash != calculateHash(b) {
																																			return false
																																					}
																																							if i > 0 && b.PreviousHash != chain[i-1].Hash {
																																										return false
																																												}
																																													}
																																														return true
																																														}

																																														func main() {
																																															genesis := createBlock(0, "Bharat Eagle Genesis", "0")
																																																block1 := createBlock(1, "BEC test transaction", genesis.Hash)

																																																	chain := []Block{genesis, block1}

																																																		fmt.Println("Bharat Eagle Chain (BEC)")
																																																			fmt.Println("Before tampering:", isChainValid(chain))

																																																				// Test: change block data without updating its hash.
																																																					chain[1].Data = "Changed data - not original"

																																																						fmt.Println("After tampering:", isChainValid(chain))
																																																						}
																																																						