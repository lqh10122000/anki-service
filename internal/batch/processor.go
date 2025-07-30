package batch

import (
	"complaint-service/internal/repository"
	"log"
)

type Processor struct {
	Repo repository.CustomerRepository
}

func (p *Processor) RunBatch() {
	log.Println("Batch processing started...")
	customers, err := p.Repo.Count()
	if err != nil {
		log.Printf("Error counting customers: %v\n", err)
		return
	}

	log.Printf("Total customers: %d\n", customers)
}
