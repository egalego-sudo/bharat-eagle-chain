				
				package main

				import (
					"crypto/sha256"
						"fmt"
							"time"
							)

							type Block struct {
								Number     int
									Timestamp  string
										Data       string
											PrevHash   string
												Hash       string
												}

												type Blockchain struct {
													Blocks []Block
													}

													func calculateHash(b Block) string {
														record := fmt.Sprintf(
																"%d|%s|%s|%s",
																		b.Number,
																				b.Timestamp,
																						b.Data,
																								b.PrevHash,
																									)

																										sum := sha256.Sum256([]byte(record))
																											return fmt.Sprintf("%x", sum)
																											}

																											func (bc *Blockchain) AddBlock(data string) {
																												var previousHash string
																													number := len(bc.Blocks)

																														if number > 0 {
																																previousHash = bc.Blocks[number-1].Hash
																																	}

																																		block := Block{
																																				Number:    number,
																																						Timestamp: time.Now().Format(time.RFC3339),
																																								Data:      data,
																																										PrevHash:  previousHash,
																																											}

																																												block.Hash = calculateHash(block)
																																													bc.Blocks = append(bc.Blocks, block)
																																													}

																																													func (bc Blockchain) IsValid() bool {
																																														for i, block := range bc.Blocks {
																																																if block.Hash != calculateHash(block) {
																																																			return false
																																																					}

																																																							if i == 0 {
																																																										if block.PrevHash != "" {
																																																														return false
																																																																	}
																																																																			} else {
																																																																						if block.PrevHash != bc.Blocks[i-1].Hash {
																																																																										return false
																																																																													}
																																																																															}
																																																																																}

																																																																																	return true
																																																																																	}

																																																																																	func main() {
																																																																																		chain := Blockchain{}

																																																																																			chain.AddBlock("BEC Genesis Block")
																																																																																				chain.AddBlock("First BEC test transaction")
																																																																																					chain.AddBlock("Second BEC test transaction")

																																																																																						fmt.Println("=== Bharat Eagle Chain (BEC) ===")
																																																																																							fmt.Println("Total blocks:", len(chain.Blocks))
																																																																																								fmt.Println()

																																																																																									for _, block := range chain.Blocks {
																																																																																											fmt.Println("Block Number:", block.Number)
																																																																																													fmt.Println("Time:", block.Timestamp)
																																																																																															fmt.Println("Data:", block.Data)
																																																																																																	fmt.Println("Previous Hash:", block.PrevHash)
																																																																																																			fmt.Println("Hash:", block.Hash)
																																																																																																					fmt.Println("------------------------------")
																																																																																																						}

																																																																																																							fmt.Println("Blockchain valid:", chain.IsValid())
																																																																																																							}

																																																																																																							