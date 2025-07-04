package lcu

type EventTopic struct{}

var Events EventTopic

// Gameflow Events
func (EventTopic) GameflowPhase() string   { return "OnJsonApiEvent_lol-gameflow_v1_gameflow-phase" }
func (EventTopic) GameflowSession() string { return "OnJsonApiEvent_lol-gameflow_v1_session" }
func (EventTopic) GameflowGameData() string {
	return "OnJsonApiEvent_lol-gameflow_v1_gameflow-metadata_game-data"
}
func (EventTopic) GameflowGameClient() string {
	return "OnJsonApiEvent_lol-gameflow_v1_gameflow-metadata_game-client"
}
func (EventTopic) GameflowRegistration() string {
	return "OnJsonApiEvent_lol-gameflow_v1_gameflow-metadata_registration"
}

// Champselect Events
func (EventTopic) ChampSelectSession() string { return "OnJsonApiEvent_lol-champ-select_v1_session" }
func (EventTopic) ChampSelectCurrentChampion() string {
	return "OnJsonApiEvent_lol-champ-select_v1_current-champion"
}
func (EventTopic) ChampSelectBenchEnabled() string {
	return "OnJsonApiEvent_lol-champ-select_v1_bench-enabled"
}
func (EventTopic) ChampSelectBenchChampions() string {
	return "OnJsonApiEvent_lol-champ-select_v1_bench-swap-list"
}
func (EventTopic) ChampSelectPickableChampions() string {
	return "OnJsonApiEvent_lol-champ-select_v1_pickable-champions"
}
func (EventTopic) ChampSelectBannableChampions() string {
	return "OnJsonApiEvent_lol-champ-select_v1_bannable-champions"
}
func (EventTopic) ChampSelectPickOrderSwaps() string {
	return "OnJsonApiEvent_lol-champ-select_v1_pick-order-swaps"
}
func (EventTopic) ChampSelectOngoingTrade() string {
	return "OnJsonApiEvent_lol-champ-select_v1_ongoing-trade"
}

// Lobby Events
func (EventTopic) Lobby() string            { return "OnJsonApiEvent_lol-lobby_v2_lobby" }
func (EventTopic) LobbyInvitations() string { return "OnJsonApiEvent_lol-lobby_v2_lobby-invitations" }
func (EventTopic) LobbyMatchmakingSearch() string {
	return "OnJsonApiEvent_lol-lobby_v2_lobby-matchmaking-search-state"
}
func (EventTopic) LobbyReceivedInvitations() string {
	return "OnJsonApiEvent_lol-lobby_v2_received-invitations"
}
func (EventTopic) LobbyEligibilityRestrictions() string {
	return "OnJsonApiEvent_lol-lobby_v2_eligibility-restrictions"
}
func (EventTopic) LobbyPartyRewards() string { return "OnJsonApiEvent_lol-lobby_v2_party-rewards" }

// Matchmaking Events
func (EventTopic) MatchmakingSearch() string { return "OnJsonApiEvent_lol-matchmaking_v1_search" }
func (EventTopic) MatchmakingReadyCheck() string {
	return "OnJsonApiEvent_lol-matchmaking_v1_ready-check"
}
func (EventTopic) MatchmakingDodgeData() string {
	return "OnJsonApiEvent_lol-matchmaking_v1_dodge-data"
}

// Summoner Events
func (EventTopic) SummonerCurrent() string { return "OnJsonApiEvent_lol-summoner_v1_current-summoner" }
func (EventTopic) SummonerStatus() string  { return "OnJsonApiEvent_lol-summoner_v1_status" }
func (EventTopic) SummonerProfile() string { return "OnJsonApiEvent_lol-summoner_v1_summoner-profile" }

// Chat Events
func (EventTopic) ChatConversations() string { return "OnJsonApiEvent_lol-chat_v1_conversations" }
func (EventTopic) ChatFriends() string       { return "OnJsonApiEvent_lol-chat_v1_friends" }
func (EventTopic) ChatMe() string            { return "OnJsonApiEvent_lol-chat_v1_me" }
func (EventTopic) ChatBlocked() string       { return "OnJsonApiEvent_lol-chat_v1_blocked-and-muted" }
func (EventTopic) ChatSettings() string      { return "OnJsonApiEvent_lol-chat_v1_settings" }
func (EventTopic) ChatSession() string       { return "OnJsonApiEvent_lol-chat_v1_session" }

// Game Client Events
func (EventTopic) GameClientChatCombatText() string {
	return "OnJsonApiEvent_lol-game-client-chat_v1_combat-text"
}
func (EventTopic) GameClientChatInstantMessages() string {
	return "OnJsonApiEvent_lol-game-client-chat_v1_instant-messages"
}

// End of Game Events
func (EventTopic) EndOfGameStatsBlock() string {
	return "OnJsonApiEvent_lol-end-of-game_v1_eog-stats-block"
}
func (EventTopic) EndOfGameChampionMastery() string {
	return "OnJsonApiEvent_lol-end-of-game_v1_champion-mastery-updates"
}
func (EventTopic) EndOfGameTFTStatsBlock() string {
	return "OnJsonApiEvent_lol-end-of-game_v1_tft-eog-stats"
}

// Ranked Events
func (EventTopic) RankedStats() string         { return "OnJsonApiEvent_lol-ranked_v1_current-ranked-stats" }
func (EventTopic) RankedNotifications() string { return "OnJsonApiEvent_lol-ranked_v1_notifications" }
func (EventTopic) RankedSigned() string        { return "OnJsonApiEvent_lol-ranked_v1_signed-ranked-stats" }

// Store Events
func (EventTopic) StoreWallet() string  { return "OnJsonApiEvent_lol-store_v1_wallet" }
func (EventTopic) StoreCatalog() string { return "OnJsonApiEvent_lol-store_v1_catalog" }
func (EventTopic) StoreOrderNotifications() string {
	return "OnJsonApiEvent_lol-store_v1_order-notifications"
}

// Loot Events
func (EventTopic) LootPlayerLoot() string    { return "OnJsonApiEvent_lol-loot_v1_player-loot" }
func (EventTopic) LootPlayerLootMap() string { return "OnJsonApiEvent_lol-loot_v1_player-loot-map" }
func (EventTopic) LootPlayerLootNotifications() string {
	return "OnJsonApiEvent_lol-loot_v1_player-loot-notifications"
}
func (EventTopic) LootPlayerDisplayCategories() string {
	return "OnJsonApiEvent_lol-loot_v1_player-display-categories"
}
func (EventTopic) LootNewPlayerLoot() string { return "OnJsonApiEvent_lol-loot_v1_new-player-loot" }
func (EventTopic) LootRecipes() string       { return "OnJsonApiEvent_lol-loot_v1_recipes" }

// Runes related Events
func (EventTopic) PerksInventory() string   { return "OnJsonApiEvent_lol-perks_v1_inventory" }
func (EventTopic) PerksPages() string       { return "OnJsonApiEvent_lol-perks_v1_pages" }
func (EventTopic) PerksCurrentPage() string { return "OnJsonApiEvent_lol-perks_v1_currentpage" }
func (EventTopic) PerksStyles() string      { return "OnJsonApiEvent_lol-perks_v1_styles" }
func (EventTopic) PerksSettings() string    { return "OnJsonApiEvent_lol-perks_v1_settings" }

// Collections Events
func (EventTopic) CollectionsChampionMastery() string {
	return "OnJsonApiEvent_lol-collections_v1_inventories_champion-mastery"
}
func (EventTopic) CollectionsWardSkins() string {
	return "OnJsonApiEvent_lol-collections_v1_inventories_ward-skins"
}
func (EventTopic) CollectionsChampionSkins() string {
	return "OnJsonApiEvent_lol-collections_v1_inventories_skins"
}
func (EventTopic) CollectionsChampions() string {
	return "OnJsonApiEvent_lol-collections_v1_inventories_champions"
}
func (EventTopic) CollectionsEmotes() string {
	return "OnJsonApiEvent_lol-collections_v1_inventories_emotes"
}
func (EventTopic) CollectionsSpells() string {
	return "OnJsonApiEvent_lol-collections_v1_inventories_spells"
}

// Challenges Events
func (EventTopic) ChallengesData() string { return "OnJsonApiEvent_lol-challenges_v1_challenges-data" }
func (EventTopic) ChallengesPersonalData() string {
	return "OnJsonApiEvent_lol-challenges_v1_challenges-personal-data"
}
func (EventTopic) ChallengesUpdatedChallenges() string {
	return "OnJsonApiEvent_lol-challenges_v1_updated-challenges"
}
func (EventTopic) ChallengesSeasonRewards() string {
	return "OnJsonApiEvent_lol-challenges_v1_season-rewards"
}
func (EventTopic) ChallengesLeveledUp() string { return "OnJsonApiEvent_lol-challenges_v1_leveled-up" }

// Replays Events
func (EventTopic) ReplaysMetadata() string      { return "OnJsonApiEvent_lol-replays_v1_metadata" }
func (EventTopic) ReplaysConfiguration() string { return "OnJsonApiEvent_lol-replays_v1_configuration" }
func (EventTopic) ReplaysRoflsPath() string     { return "OnJsonApiEvent_lol-replays_v1_rofls-path" }

// Spectator Events
func (EventTopic) SpectatorDelayedObserverV1() string {
	return "OnJsonApiEvent_lol-spectator_v1_delayed-observer-v1"
}
func (EventTopic) SpectatorSessionToken() string {
	return "OnJsonApiEvent_lol-spectator_v1_session-token"
}

// Settings Events
func (EventTopic) SettingsLocal() string   { return "OnJsonApiEvent_lol-settings_v1_local" }
func (EventTopic) SettingsReady() string   { return "OnJsonApiEvent_lol-settings_v1_ready" }
func (EventTopic) SettingsAccount() string { return "OnJsonApiEvent_lol-settings_v2_account" }
func (EventTopic) SettingsHotkeys() string { return "OnJsonApiEvent_lol-settings_v1_hotkeys" }

// Clash Events
func (EventTopic) ClashCurrentTournament() string {
	return "OnJsonApiEvent_lol-clash_v1_currentTournament"
}
func (EventTopic) ClashEvents() string        { return "OnJsonApiEvent_lol-clash_v1_events" }
func (EventTopic) ClashPlayerData() string    { return "OnJsonApiEvent_lol-clash_v1_player-data" }
func (EventTopic) ClashPlayerHistory() string { return "OnJsonApiEvent_lol-clash_v1_player-history" }
func (EventTopic) ClashPlayerRewards() string { return "OnJsonApiEvent_lol-clash_v1_player-rewards" }
func (EventTopic) ClashRoster() string        { return "OnJsonApiEvent_lol-clash_v1_roster" }
func (EventTopic) ClashSimpleStateFlags() string {
	return "OnJsonApiEvent_lol-clash_v1_simple-state-flags"
}
func (EventTopic) ClashTournamentSummary() string {
	return "OnJsonApiEvent_lol-clash_v1_tournament-summary"
}
func (EventTopic) ClashTournamentStateInfo() string {
	return "OnJsonApiEvent_lol-clash_v1_tournament-state-info"
}
func (EventTopic) ClashVisibility() string { return "OnJsonApiEvent_lol-clash_v1_visibility" }
func (EventTopic) ClashEogPlayerUpdate() string {
	return "OnJsonApiEvent_lol-clash_v1_eog-player-update"
}

// Missions Events
func (EventTopic) MissionsPlayerMissions() string {
	return "OnJsonApiEvent_lol-missions_v1_player-missions"
}
func (EventTopic) MissionsPlayerMissionsComplete() string {
	return "OnJsonApiEvent_lol-missions_v1_player-missions-complete"
}
func (EventTopic) MissionsPlayerMissionsProgress() string {
	return "OnJsonApiEvent_lol-missions_v1_player-missions-progress"
}
func (EventTopic) MissionsPlayerMissionsOverview() string {
	return "OnJsonApiEvent_lol-missions_v1_player-missions-overview"
}

// Patch Events
func (EventTopic) PatcherProductState() string {
	return "OnJsonApiEvent_patcher_v1_products_league_of_legends_state"
}
func (EventTopic) PatcherProductInstallLocation() string {
	return "OnJsonApiEvent_patcher_v1_products_league_of_legends_install-location"
}
func (EventTopic) PatcherProductDetectedCorruption() string {
	return "OnJsonApiEvent_patcher_v1_products_league_of_legends_detected-corruption-type"
}
func (EventTopic) PatcherNotifications() string { return "OnJsonApiEvent_patcher_v1_notifications" }
func (EventTopic) PatcherStatus() string        { return "OnJsonApiEvent_patcher_v1_status" }
func (EventTopic) PatcherP2PStatus() string     { return "OnJsonApiEvent_patcher_v1_p2p-status" }
func (EventTopic) PatcherP2PStatusV2() string   { return "OnJsonApiEvent_patcher_v1_p2p-status-v2" }
func (EventTopic) PatcherEnvironment() string   { return "OnJsonApiEvent_patcher_v1_environment" }

// System Events
func (EventTopic) SystemBuildsInstall() string   { return "OnJsonApiEvent_system_v1_builds_install" }
func (EventTopic) SystemBuildsUninstall() string { return "OnJsonApiEvent_system_v1_builds_uninstall" }

// Plugin Manager Events
func (EventTopic) PluginManagerDescription() string {
	return "OnJsonApiEvent_plugin-manager_v1_external-plugins_description"
}
func (EventTopic) PluginManagerStatus() string  { return "OnJsonApiEvent_plugin-manager_v1_status" }
func (EventTopic) PluginManagerPlugins() string { return "OnJsonApiEvent_plugin-manager_v1_plugins" }
func (EventTopic) PluginManagerExternalPlugins() string {
	return "OnJsonApiEvent_plugin-manager_v1_external-plugins_availability"
}

// Process Control Events
func (EventTopic) ProcessControlV1() string { return "OnJsonApiEvent_process-control_v1_process" }

// Installation Events
func (EventTopic) InstallationV1() string { return "OnJsonApiEvent_installation_v1_products-installed" }

// Service Proxy Events
func (EventTopic) ServiceProxyV1() string { return "OnJsonApiEvent_service-proxy_v1_auth-token" }

// RMS Events
func (EventTopic) RmsV1() string { return "OnJsonApiEvent_rms_v1_champion-select-player-selection" }

// Login Events
func (EventTopic) LoginV1() string { return "OnJsonApiEvent_login_v1_login-platform-credentials" }

// Entitlements Events
func (EventTopic) EntitlementsV1() string { return "OnJsonApiEvent_entitlements_v1_token" }

// Honor Events
func (EventTopic) HonorV2Ballot() string { return "OnJsonApiEvent_lol-honor-v2_v1_ballot" }
func (EventTopic) HonorV2LateRecognition() string {
	return "OnJsonApiEvent_lol-honor-v2_v1_late-recognition"
}
func (EventTopic) HonorV2Mutual() string      { return "OnJsonApiEvent_lol-honor-v2_v1_mutual-honor" }
func (EventTopic) HonorV2Profile() string     { return "OnJsonApiEvent_lol-honor-v2_v1_profile" }
func (EventTopic) HonorV2Recognition() string { return "OnJsonApiEvent_lol-honor-v2_v1_recognition" }
func (EventTopic) HonorV2Reward() string      { return "OnJsonApiEvent_lol-honor-v2_v1_reward-granted" }
func (EventTopic) HonorV2TeamChoices() string { return "OnJsonApiEvent_lol-honor-v2_v1_team-choices" }
func (EventTopic) HonorV2VoteCompletion() string {
	return "OnJsonApiEvent_lol-honor-v2_v1_vote-completion"
}

// TFT Events
func (EventTopic) TftDamageAnalysis() string { return "OnJsonApiEvent_lol-tft_v1_damage-analysis" }
func (EventTopic) TftPasses() string         { return "OnJsonApiEvent_lol-tft_v1_passes" }
func (EventTopic) TftBattlepass() string     { return "OnJsonApiEvent_lol-tft_v1_battlepass" }

// Career Stats Events
func (EventTopic) CareerStatsV1() string {
	return "OnJsonApiEvent_lol-career-stats_v1_champion-averages"
}
func (EventTopic) CareerStatsPosition() string {
	return "OnJsonApiEvent_lol-career-stats_v1_position-averages"
}
func (EventTopic) CareerStatsSummoner() string {
	return "OnJsonApiEvent_lol-career-stats_v1_summoner-stats"
}

// Item Sets Events
func (EventTopic) ItemSetsV1() string { return "OnJsonApiEvent_lol-item-sets_v1_item-sets" }

// Loadouts Events
func (EventTopic) LoadoutsV4() string { return "OnJsonApiEvent_lol-loadouts_v4_loadouts" }

// Account Events
func (EventTopic) AccountV1() string { return "OnJsonApiEvent_lol-account_v1_accounts" }

// Generic/Callback Events
func (EventTopic) OnCallback(callback string) string { return "OnCallback_" + callback }

// LCDS Events
func (EventTopic) OnLcdsEvent(event string) string { return "OnLcdsEvent_" + event }

// All Events
func (EventTopic) All() string { return "OnJsonApiEvent" }
