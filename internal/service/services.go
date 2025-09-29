package service

type Services struct {
	providerService *SPService
	historyService  *HistoryService
	sessionService  *SessionService
	shortcutService *ShortcutService
}

func NewServices(
	providersService *SPService,
	historyService *HistoryService,
	sessionService *SessionService,
	shortcutService *ShortcutService,
) *Services {
	return &Services{
		providerService: providersService,
		historyService:  historyService,
		sessionService:  sessionService,
		shortcutService: shortcutService,
	}
}

func (service *Services) GetProvidersService() *SPService {
	return service.providerService
}

func (service *Services) GetHistoryService() *HistoryService {
	return service.historyService
}

func (service *Services) GetSessionService() *SessionService {
	return service.sessionService
}

func (service *Services) GetShortcutService() *ShortcutService {
	return service.shortcutService
}
