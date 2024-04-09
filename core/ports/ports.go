package ports

import "github.com/Pokemon-Infinite-Stadium/database-updater/core/domain"

type EntryAdapter interface {
}

type ExitAdapter interface {
	SavePokemons(pokemons []domain.Pokemon) error
}
