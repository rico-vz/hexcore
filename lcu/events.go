package lcu

type EventTopic struct{}

var Events EventTopic

// Fired when a callback is invoked.
func (EventTopic) Callback() string { return "OnCallback" }

// Fired when a resource is changed.
func (EventTopic) JsonApiEvent() string { return "OnJsonApiEvent" }

// Fired when a message is received over the LCDS connection.
func (EventTopic) LcdsEvent() string { return "OnLcdsEvent" }

// Fired when a entry of sufficient severity is logged.
func (EventTopic) Log() string { return "OnLog" }

// Chat events
func (EventTopic) ChatV1Session() string        { return "OnJsonApiEvent_chat_v1_session" }
func (EventTopic) ChatV1Settings() string       { return "OnJsonApiEvent_chat_v1_settings" }
func (EventTopic) ChatV3Blocked() string        { return "OnJsonApiEvent_chat_v3_blocked" }
func (EventTopic) ChatV3Errors() string         { return "OnJsonApiEvent_chat_v3_errors" }
func (EventTopic) ChatV3Friends() string        { return "OnJsonApiEvent_chat_v3_friends" }
func (EventTopic) ChatV3Groups() string         { return "OnJsonApiEvent_chat_v3_groups" }
func (EventTopic) ChatV4Presences() string      { return "OnJsonApiEvent_chat_v4_presences" }
func (EventTopic) ChatV5Messages() string       { return "OnJsonApiEvent_chat_v5_messages" }
func (EventTopic) ChatV5Participants() string   { return "OnJsonApiEvent_chat_v5_participants" }
func (EventTopic) ChatV6Conversations() string  { return "OnJsonApiEvent_chat_v6_conversations" }
func (EventTopic) ChatV6Friendrequests() string { return "OnJsonApiEvent_chat_v6_friendrequests" }

// Client Config events
func (EventTopic) ClientConfigV1Status() string { return "OnJsonApiEvent_client-config_v1_status" }
func (EventTopic) ClientConfigV2Config() string { return "OnJsonApiEvent_client-config_v2_config" }
func (EventTopic) ClientConfigV2Namespace() string {
	return "OnJsonApiEvent_client-config_v2_namespace"
}

// Data Store events
func (EventTopic) DataStoreV1InstallSettings() string {
	return "OnJsonApiEvent_data-store_v1_install-settings"
}
func (EventTopic) DataStoreV1SystemSettings() string {
	return "OnJsonApiEvent_data-store_v1_system-settings"
}

// Deep Links events
func (EventTopic) DeepLinksV1Settings() string { return "OnJsonApiEvent_deep-links_v1_settings" }

// Entitlements events
func (EventTopic) EntitlementsV1Token() string { return "OnJsonApiEvent_entitlements_v1_token" }
func (EventTopic) EntitlementsV2Token() string { return "OnJsonApiEvent_entitlements_v2_token" }

// Ga Restriction events
func (EventTopic) GaRestrictionV1PenaltyNotifications() string {
	return "OnJsonApiEvent_ga-restriction_v1_penalty-notifications"
}

// Lol Account Verification events
func (EventTopic) AccountVerificationV1IsVerified() string {
	return "OnJsonApiEvent_lol-account-verification_v1_is-verified"
}

// Lol Active Boosts events
func (EventTopic) ActiveBoostsV1() string { return "OnJsonApiEvent_lol-active-boosts_v1_active-boosts" }

// Lol Activity Center events
func (EventTopic) ActivityCenterV1Ready() string {
	return "OnJsonApiEvent_lol-activity-center_v1_ready"
}

// Lol Anti Addiction events
func (EventTopic) AntiAddictionV1Token() string {
	return "OnJsonApiEvent_lol-anti-addiction_v1_anti-addiction-token"
}

// Lol Cap Missions events
func (EventTopic) CapMissionsV1Ready() string { return "OnJsonApiEvent_lol-cap-missions_v1_ready" }

// Lol Challenges events
func (EventTopic) ChallengesV1ClientState() string {
	return "OnJsonApiEvent_lol-challenges_v1_client-state"
}
func (EventTopic) ChallengesV1Seasons() string { return "OnJsonApiEvent_lol-challenges_v1_seasons" }

// Lol Champ Select events
func (EventTopic) ChampSelectV1AllGridChampions() string {
	return "OnJsonApiEvent_lol-champ-select_v1_all-grid-champions"
}
func (EventTopic) ChampSelectV1GridChampions() string {
	return "OnJsonApiEvent_lol-champ-select_v1_grid-champions"
}
func (EventTopic) ChampSelectV1MutedPlayers() string {
	return "OnJsonApiEvent_lol-champ-select_v1_muted-players"
}
func (EventTopic) ChampSelectV1Session() string { return "OnJsonApiEvent_lol-champ-select_v1_session" }
func (EventTopic) ChampSelectV1TeamBoost() string {
	return "OnJsonApiEvent_lol-champ-select_v1_team-boost"
}

// Lol Champ Select Legacy events
func (EventTopic) ChampSelectLegacyV1BannableChampionIds() string {
	return "OnJsonApiEvent_lol-champ-select-legacy_v1_bannable-champion-ids"
}
func (EventTopic) ChampSelectLegacyV1CurrentChampion() string {
	return "OnJsonApiEvent_lol-champ-select-legacy_v1_current-champion"
}
func (EventTopic) ChampSelectLegacyV1DisabledChampionIds() string {
	return "OnJsonApiEvent_lol-champ-select-legacy_v1_disabled-champion-ids"
}
func (EventTopic) ChampSelectLegacyV1ImplementationActive() string {
	return "OnJsonApiEvent_lol-champ-select-legacy_v1_implementation-active"
}
func (EventTopic) ChampSelectLegacyV1PickableChampionIds() string {
	return "OnJsonApiEvent_lol-champ-select-legacy_v1_pickable-champion-ids"
}
func (EventTopic) ChampSelectLegacyV1PickableSkinIds() string {
	return "OnJsonApiEvent_lol-champ-select-legacy_v1_pickable-skin-ids"
}
func (EventTopic) ChampSelectLegacyV1Session() string {
	return "OnJsonApiEvent_lol-champ-select-legacy_v1_session"
}
func (EventTopic) ChampSelectLegacyV1TeamBoost() string {
	return "OnJsonApiEvent_lol-champ-select-legacy_v1_team-boost"
}

// Lol Champion Mastery events
func (EventTopic) ChampionMasteryV1LocalPlayer() string {
	return "OnJsonApiEvent_lol-champion-mastery_v1_local-player"
}
func (EventTopic) ChampionMasteryV1Ready() string {
	return "OnJsonApiEvent_lol-champion-mastery_v1_ready"
}

// Lol Champions events
func (EventTopic) ChampionsV1Inventories() string {
	return "OnJsonApiEvent_lol-champions_v1_inventories"
}
func (EventTopic) ChampionsV1OwnedChampionsMinimal() string {
	return "OnJsonApiEvent_lol-champions_v1_owned-champions-minimal"
}

// Lol Chat events
func (EventTopic) LoLChatV1BlockedPlayers() string {
	return "OnJsonApiEvent_lol-chat_v1_blocked-players"
}
func (EventTopic) LoLChatV1Config() string        { return "OnJsonApiEvent_lol-chat_v1_config" }
func (EventTopic) LoLChatV1Conversations() string { return "OnJsonApiEvent_lol-chat_v1_conversations" }
func (EventTopic) LoLChatV1FriendCounts() string  { return "OnJsonApiEvent_lol-chat_v1_friend-counts" }
func (EventTopic) LoLChatV1FriendGroups() string  { return "OnJsonApiEvent_lol-chat_v1_friend-groups" }
func (EventTopic) LoLChatV1FriendRequests() string {
	return "OnJsonApiEvent_lol-chat_v1_friend-requests"
}
func (EventTopic) LoLChatV1Friends() string   { return "OnJsonApiEvent_lol-chat_v1_friends" }
func (EventTopic) LoLChatV1Me() string        { return "OnJsonApiEvent_lol-chat_v1_me" }
func (EventTopic) LoLChatV1Resources() string { return "OnJsonApiEvent_lol-chat_v1_resources" }
func (EventTopic) LoLChatV1Session() string   { return "OnJsonApiEvent_lol-chat_v1_session" }
func (EventTopic) LoLChatV1Settings() string  { return "OnJsonApiEvent_lol-chat_v1_settings" }
func (EventTopic) LoLChatV2FriendRequests() string {
	return "OnJsonApiEvent_lol-chat_v2_friend-requests"
}

// Lol Clash events
func (EventTopic) ClashV1CheckinAllowed() string {
	return "OnJsonApiEvent_lol-clash_v1_checkin-allowed"
}
func (EventTopic) ClashV1DisabledConfig() string {
	return "OnJsonApiEvent_lol-clash_v1_disabled-config"
}
func (EventTopic) ClashV1Enabled() string    { return "OnJsonApiEvent_lol-clash_v1_enabled" }
func (EventTopic) ClashV1Iconconfig() string { return "OnJsonApiEvent_lol-clash_v1_iconconfig" }
func (EventTopic) ClashV1InvitedRosterIds() string {
	return "OnJsonApiEvent_lol-clash_v1_invited-roster-ids"
}
func (EventTopic) ClashV1Player() string { return "OnJsonApiEvent_lol-clash_v1_player" }
func (EventTopic) ClashV1PlaymodeRestricted() string {
	return "OnJsonApiEvent_lol-clash_v1_playmode-restricted"
}
func (EventTopic) ClashV1Ready() string { return "OnJsonApiEvent_lol-clash_v1_ready" }
func (EventTopic) ClashV1SimpleStateFlags() string {
	return "OnJsonApiEvent_lol-clash_v1_simple-state-flags"
}
func (EventTopic) ClashV1Time() string         { return "OnJsonApiEvent_lol-clash_v1_time" }
func (EventTopic) ClashV1Visible() string      { return "OnJsonApiEvent_lol-clash_v1_visible" }
func (EventTopic) ClashV1VoiceEnabled() string { return "OnJsonApiEvent_lol-clash_v1_voice-enabled" }
func (EventTopic) ClashV2PlaymodeRestricted() string {
	return "OnJsonApiEvent_lol-clash_v2_playmode-restricted"
}

// Lol Client Config events
func (EventTopic) ClientConfigV3() string { return "OnJsonApiEvent_lol-client-config_v3_client-config" }

// Lol Collections events
func (EventTopic) CollectionsV1Inventories() string {
	return "OnJsonApiEvent_lol-collections_v1_inventories"
}

// Lol Cosmetics events
func (EventTopic) CosmeticsV1Favorites() string { return "OnJsonApiEvent_lol-cosmetics_v1_favorites" }
func (EventTopic) CosmeticsV1Inventories() string {
	return "OnJsonApiEvent_lol-cosmetics_v1_inventories"
}

// Lol Drops events
func (EventTopic) DropsV1Ready() string { return "OnJsonApiEvent_lol-drops_v1_ready" }

// Lol Dx9 Deprecation events
func (EventTopic) Dx9DeprecationNeedsHardwareUpgrade() string {
	return "OnJsonApiEvent_lol-dx9-deprecation_needs-hardware-upgrade"
}

// Lol End Of Game events
func (EventTopic) EndOfGameV1EogStatsBlock() string {
	return "OnJsonApiEvent_lol-end-of-game_v1_eog-stats-block"
}
func (EventTopic) EndOfGameV1GameclientEogStatsBlock() string {
	return "OnJsonApiEvent_lol-end-of-game_v1_gameclient-eog-stats-block"
}
func (EventTopic) EndOfGameV1TftEogStats() string {
	return "OnJsonApiEvent_lol-end-of-game_v1_tft-eog-stats"
}

// Lol Event Hub events
func (EventTopic) EventHubV1Events() string { return "OnJsonApiEvent_lol-event-hub_v1_events" }
func (EventTopic) EventHubV1NavigationButtonData() string {
	return "OnJsonApiEvent_lol-event-hub_v1_navigation-button-data"
}
func (EventTopic) EventHubV1Skins() string { return "OnJsonApiEvent_lol-event-hub_v1_skins" }
func (EventTopic) EventHubV1TokenUpsell() string {
	return "OnJsonApiEvent_lol-event-hub_v1_token-upsell"
}

// Lol Event Mission events
func (EventTopic) EventMissionV1() string { return "OnJsonApiEvent_lol-event-mission_v1_event-mission" }

// Lol Game Client Chat events
func (EventTopic) GameClientChatV1Buddies() string {
	return "OnJsonApiEvent_lol-game-client-chat_v1_buddies"
}
func (EventTopic) GameClientChatV1InstantMessages() string {
	return "OnJsonApiEvent_lol-game-client-chat_v1_instant-messages"
}
func (EventTopic) GameClientChatV2Buddies() string {
	return "OnJsonApiEvent_lol-game-client-chat_v2_buddies"
}
func (EventTopic) GameClientChatV2InstantMessages() string {
	return "OnJsonApiEvent_lol-game-client-chat_v2_instant-messages"
}
func (EventTopic) GameClientChatV2PlaytimeReminderHoursPlayed() string {
	return "OnJsonApiEvent_lol-game-client-chat_v2_playtime-reminder-hours-played"
}

// Lol Game Data Inventory events
func (EventTopic) GameDataInventoryV1Ready() string {
	return "OnJsonApiEvent_lol-game-data-inventory_v1_ready"
}

// Lol Game Queues events
func (EventTopic) GameQueuesV1Custom() string { return "OnJsonApiEvent_lol-game-queues_v1_custom" }
func (EventTopic) GameQueuesV1CustomNonDefault() string {
	return "OnJsonApiEvent_lol-game-queues_v1_custom-non-default"
}
func (EventTopic) GameQueuesV1Queues() string { return "OnJsonApiEvent_lol-game-queues_v1_queues" }

// Lol Game Settings events
func (EventTopic) GameSettingsV1() string { return "OnJsonApiEvent_lol-game-settings_v1_game-settings" }
func (EventTopic) GameSettingsV1InputSettings() string {
	return "OnJsonApiEvent_lol-game-settings_v1_input-settings"
}
func (EventTopic) GameSettingsV1Ready() string { return "OnJsonApiEvent_lol-game-settings_v1_ready" }

// Lol Gameflow events
func (EventTopic) GameflowV1Availability() string {
	return "OnJsonApiEvent_lol-gameflow_v1_availability"
}
func (EventTopic) GameflowV1BattleTraining() string {
	return "OnJsonApiEvent_lol-gameflow_v1_battle-training"
}
func (EventTopic) GameflowV1EarlyExitNotifications() string {
	return "OnJsonApiEvent_lol-gameflow_v1_early-exit-notifications"
}
func (EventTopic) GameflowV1Metadata() string {
	return "OnJsonApiEvent_lol-gameflow_v1_gameflow-metadata"
}
func (EventTopic) GameflowV1Phase() string    { return "OnJsonApiEvent_lol-gameflow_v1_gameflow-phase" }
func (EventTopic) GameflowV1Session() string  { return "OnJsonApiEvent_lol-gameflow_v1_session" }
func (EventTopic) GameflowV1Spectate() string { return "OnJsonApiEvent_lol-gameflow_v1_spectate" }
func (EventTopic) GameflowV1Watch() string    { return "OnJsonApiEvent_lol-gameflow_v1_watch" }

// Lol Highlights events
func (EventTopic) HighlightsV1Config() string { return "OnJsonApiEvent_lol-highlights_v1_config" }
func (EventTopic) HighlightsV1FolderPath() string {
	return "OnJsonApiEvent_lol-highlights_v1_highlights-folder-path"
}

// Lol Honeyfruit events
func (EventTopic) HoneyfruitV1VngPublisherSettings() string {
	return "OnJsonApiEvent_lol-honeyfruit_v1_vng-publisher-settings"
}

// Lol Honor V2 events
func (EventTopic) HonorV2V1Ballot() string      { return "OnJsonApiEvent_lol-honor-v2_v1_ballot" }
func (EventTopic) HonorV2V1Config() string      { return "OnJsonApiEvent_lol-honor-v2_v1_config" }
func (EventTopic) HonorV2V1LevelChange() string { return "OnJsonApiEvent_lol-honor-v2_v1_level-change" }
func (EventTopic) HonorV2V1Profile() string     { return "OnJsonApiEvent_lol-honor-v2_v1_profile" }
func (EventTopic) HonorV2V1Recipients() string  { return "OnJsonApiEvent_lol-honor-v2_v1_recipients" }
func (EventTopic) HonorV2V1Recognition() string { return "OnJsonApiEvent_lol-honor-v2_v1_recognition" }
func (EventTopic) HonorV2V1RecognitionHistory() string {
	return "OnJsonApiEvent_lol-honor-v2_v1_recognition-history"
}
func (EventTopic) HonorV2V1TeamChoices() string { return "OnJsonApiEvent_lol-honor-v2_v1_team-choices" }

// Lol Hovercard events
func (EventTopic) HovercardV1FriendInfo() string {
	return "OnJsonApiEvent_lol-hovercard_v1_friend-info"
}
func (EventTopic) HovercardV1FriendInfoBySummoner() string {
	return "OnJsonApiEvent_lol-hovercard_v1_friend-info-by-summoner"
}

// Lol Inventory events
func (EventTopic) InventoryV1InitialConfigurationComplete() string {
	return "OnJsonApiEvent_lol-inventory_v1_initial-configuration-complete"
}
func (EventTopic) InventoryV1() string { return "OnJsonApiEvent_lol-inventory_v1_inventory" }
func (EventTopic) InventoryV1Signedinventory() string {
	return "OnJsonApiEvent_lol-inventory_v1_signedInventory"
}
func (EventTopic) InventoryV1Wallet() string { return "OnJsonApiEvent_lol-inventory_v1_wallet" }
func (EventTopic) InventoryV2() string       { return "OnJsonApiEvent_lol-inventory_v2_inventory" }

// Lol Kr Playtime Reminder events
func (EventTopic) KrPlaytimeReminderV1HoursPlayed() string {
	return "OnJsonApiEvent_lol-kr-playtime-reminder_v1_hours-played"
}

// Lol Leaderboard events
func (EventTopic) LeaderboardV1Ready() string { return "OnJsonApiEvent_lol-leaderboard_v1_ready" }

// Lol League Session events
func (EventTopic) LeagueSessionV1Token() string {
	return "OnJsonApiEvent_lol-league-session_v1_league-session-token"
}

// Lol Leaver Buster events
func (EventTopic) LeaverBusterV1RankedRestriction() string {
	return "OnJsonApiEvent_lol-leaver-buster_v1_ranked-restriction"
}

// Lol License Agreement events
func (EventTopic) LicenseAgreementV1Agreements() string {
	return "OnJsonApiEvent_lol-license-agreement_v1_agreements"
}
func (EventTopic) LicenseAgreementV1AllAgreements() string {
	return "OnJsonApiEvent_lol-license-agreement_v1_all-agreements"
}
func (EventTopic) LicenseAgreementV1ServeLocation() string {
	return "OnJsonApiEvent_lol-license-agreement_v1_serve-location"
}

// Lol Loadouts events
func (EventTopic) LoadoutsV1Enabled() string { return "OnJsonApiEvent_lol-loadouts_v1_enabled" }
func (EventTopic) LoadoutsV1Ready() string   { return "OnJsonApiEvent_lol-loadouts_v1_loadouts-ready" }
func (EventTopic) LoadoutsV4Loadout() string { return "OnJsonApiEvent_lol-loadouts_v4_loadout" }
func (EventTopic) LoadoutsV4() string        { return "OnJsonApiEvent_lol-loadouts_v4_loadouts" }

// Lol Lobby events
func (EventTopic) LobbyV1() string            { return "OnJsonApiEvent_lol-lobby_v1_lobby" }
func (EventTopic) LobbyV2Comms() string       { return "OnJsonApiEvent_lol-lobby_v2_comms" }
func (EventTopic) LobbyV2Eligibility() string { return "OnJsonApiEvent_lol-lobby_v2_eligibility" }
func (EventTopic) LobbyV2() string            { return "OnJsonApiEvent_lol-lobby_v2_lobby" }

// Lol Lobby Team Builder events
func (EventTopic) LobbyTeamBuilderV1ChampSelect() string {
	return "OnJsonApiEvent_lol-lobby-team-builder_champ-select_v1"
}
func (EventTopic) LobbyTeamBuilderV1Matchmaking() string {
	return "OnJsonApiEvent_lol-lobby-team-builder_v1_matchmaking"
}

// Lol Lock And Load events
func (EventTopic) LockAndLoadV1HomeHubsWaits() string {
	return "OnJsonApiEvent_lol-lock-and-load_v1_home-hubs-waits"
}
func (EventTopic) LockAndLoadV1ShouldShowProgressBarText() string {
	return "OnJsonApiEvent_lol-lock-and-load_v1_should-show-progress-bar-text"
}
func (EventTopic) LockAndLoadV1ShouldWaitForHomeHubs() string {
	return "OnJsonApiEvent_lol-lock-and-load_v1_should-wait-for-home-hubs"
}

// Lol Loot events
func (EventTopic) LootV1CurrencyConfiguration() string {
	return "OnJsonApiEvent_lol-loot_v1_currency-configuration"
}
func (EventTopic) LootV1Enabled() string    { return "OnJsonApiEvent_lol-loot_v1_enabled" }
func (EventTopic) LootV1Milestones() string { return "OnJsonApiEvent_lol-loot_v1_milestones" }
func (EventTopic) LootV1Ready() string      { return "OnJsonApiEvent_lol-loot_v1_ready" }
func (EventTopic) LootV1Recipes() string    { return "OnJsonApiEvent_lol-loot_v1_recipes" }

// Lol Loyalty events
func (EventTopic) LoyaltyV1StatusNotification() string {
	return "OnJsonApiEvent_lol-loyalty_v1_status-notification"
}

// Lol Maps events
func (EventTopic) MapsV1() string { return "OnJsonApiEvent_lol-maps_v1_maps" }
func (EventTopic) MapsV2() string { return "OnJsonApiEvent_lol-maps_v2_maps" }

// Lol Marketing Preferences events
func (EventTopic) MarketingPreferencesV1Ready() string {
	return "OnJsonApiEvent_lol-marketing-preferences_v1_ready"
}

// Lol Marketplace events
func (EventTopic) MarketplaceV1Products() string { return "OnJsonApiEvent_lol-marketplace_v1_products" }
func (EventTopic) MarketplaceV1Ready() string    { return "OnJsonApiEvent_lol-marketplace_v1_ready" }

// Lol Matchmaking events
func (EventTopic) MatchmakingV1Search() string { return "OnJsonApiEvent_lol-matchmaking_v1_search" }

// Lol Missions events
func (EventTopic) MissionsV1() string       { return "OnJsonApiEvent_lol-missions_v1_missions" }
func (EventTopic) MissionsV1Series() string { return "OnJsonApiEvent_lol-missions_v1_series" }

// Lol Nacho events
func (EventTopic) NachoV1Banner() string     { return "OnJsonApiEvent_lol-nacho_v1_banner" }
func (EventTopic) NachoV1BannerOdds() string { return "OnJsonApiEvent_lol-nacho_v1_banner-odds" }
func (EventTopic) NachoV1Banners() string    { return "OnJsonApiEvent_lol-nacho_v1_banners" }
func (EventTopic) NachoV1GetActiveStoreCatalog() string {
	return "OnJsonApiEvent_lol-nacho_v1_get-active-store-catalog"
}
func (EventTopic) NachoV1GetCurrentCatalogItem() string {
	return "OnJsonApiEvent_lol-nacho_v1_get-current-catalog-item"
}
func (EventTopic) NachoV1Ready() string { return "OnJsonApiEvent_lol-nacho_v1_ready" }
func (EventTopic) NachoV1RollPurchaseEnabled() string {
	return "OnJsonApiEvent_lol-nacho_v1_roll-purchase-enabled"
}

// Lol Npe Tutorial Path events
func (EventTopic) NpeTutorialPathV1Rewards() string {
	return "OnJsonApiEvent_lol-npe-tutorial-path_v1_rewards"
}
func (EventTopic) NpeTutorialPathV1Settings() string {
	return "OnJsonApiEvent_lol-npe-tutorial-path_v1_settings"
}
func (EventTopic) NpeTutorialPathV1Tutorials() string {
	return "OnJsonApiEvent_lol-npe-tutorial-path_v1_tutorials"
}

// Lol Objectives events
func (EventTopic) ObjectivesV1() string      { return "OnJsonApiEvent_lol-objectives_v1_objectives" }
func (EventTopic) ObjectivesV1Ready() string { return "OnJsonApiEvent_lol-objectives_v1_ready" }

// Lol Patch events
func (EventTopic) PatchV1CheckingEnabled() string {
	return "OnJsonApiEvent_lol-patch_v1_checking-enabled"
}
func (EventTopic) PatchV1Environment() string   { return "OnJsonApiEvent_lol-patch_v1_environment" }
func (EventTopic) PatchV1GameVersion() string   { return "OnJsonApiEvent_lol-patch_v1_game-version" }
func (EventTopic) PatchV1Notifications() string { return "OnJsonApiEvent_lol-patch_v1_notifications" }
func (EventTopic) PatchV1ProductIntegration() string {
	return "OnJsonApiEvent_lol-patch_v1_product-integration"
}
func (EventTopic) PatchV1Products() string { return "OnJsonApiEvent_lol-patch_v1_products" }
func (EventTopic) PatchV1Status() string   { return "OnJsonApiEvent_lol-patch_v1_status" }

// Lol Perks events
func (EventTopic) PerksV1Currentpage() string { return "OnJsonApiEvent_lol-perks_v1_currentpage" }
func (EventTopic) PerksV1Inventory() string   { return "OnJsonApiEvent_lol-perks_v1_inventory" }
func (EventTopic) PerksV1Pages() string       { return "OnJsonApiEvent_lol-perks_v1_pages" }
func (EventTopic) PerksV1() string            { return "OnJsonApiEvent_lol-perks_v1_perks" }
func (EventTopic) PerksV1RuneRecommenderAutoSelect() string {
	return "OnJsonApiEvent_lol-perks_v1_rune-recommender-auto-select"
}
func (EventTopic) PerksV1Settings() string { return "OnJsonApiEvent_lol-perks_v1_settings" }
func (EventTopic) PerksV1Styles() string   { return "OnJsonApiEvent_lol-perks_v1_styles" }

// Lol Platform Config events
func (EventTopic) PlatformConfigV1InitialConfigurationComplete() string {
	return "OnJsonApiEvent_lol-platform-config_v1_initial-configuration-complete"
}
func (EventTopic) PlatformConfigV1NamespacePartiesPublishpresencedelay() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespace_Parties_PublishPresenceDelay"
}
func (EventTopic) PlatformConfigV1Namespaces() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverification() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverificationEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverificationIsnewavsserviceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification_IsNewAvsServiceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverificationKrphonedisplayenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification_KrPhoneDisplayEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverificationPasswordenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification_PasswordEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverificationPhonevalidationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification_PhoneValidationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverificationSettingsenable() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification_SettingsEnable"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverificationSettingsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification_SettingsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverificationSettingsverifyenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification_SettingsVerifyEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesAccountverificationShouldusenewavs() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_AccountVerification_ShouldUseNewAvs"
}
func (EventTopic) PlatformConfigV1NamespacesBanners() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Banners"
}
func (EventTopic) PlatformConfigV1NamespacesBannersIsenabledonprofile() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Banners_IsEnabledOnProfile"
}
func (EventTopic) PlatformConfigV1NamespacesBannersIsequipenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Banners_IsEquipEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesBannersIsothersummonersprofileenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Banners_IsOtherSummonersProfileEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesBetterrewards() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_BetterRewards"
}
func (EventTopic) PlatformConfigV1NamespacesBetterrewardsDayonemodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_BetterRewards_DayOneModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesBetterrewardsGetpostgamexpfromrms() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_BetterRewards_GetPostgameXpFromRms"
}
func (EventTopic) PlatformConfigV1NamespacesBetterrewardsVisualupdateenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_BetterRewards_VisualUpdateEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesBotconfigurations() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_BotConfigurations"
}
func (EventTopic) PlatformConfigV1NamespacesBotconfigurationsIntermediateincustoms() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_BotConfigurations_IntermediateInCustoms"
}
func (EventTopic) PlatformConfigV1NamespacesBotconfigurationsRiotscriptincustoms() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_BotConfigurations_RiotscriptInCustoms"
}
func (EventTopic) PlatformConfigV1NamespacesCareerstats() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_CareerStats"
}
func (EventTopic) PlatformConfigV1NamespacesCareerstatsStatsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_CareerStats_StatsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesCareerstatsYearsallowedtoqueryseasondata() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_CareerStats_YearsAllowedToQuerySeasonData"
}
func (EventTopic) PlatformConfigV1NamespacesChallenges() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesChallengeupdatedelayseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_ChallengeUpdateDelaySeconds"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesClientstate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_ClientState"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesCollectionenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_CollectionEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesCustomizeidentityenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_CustomizeIdentityEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesDarkmodeallowlistonly() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_DarkModeAllowlistOnly"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesEnabledinventorytypes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_EnabledInventoryTypes"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesFeatureintroenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_FeatureIntroEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesLobbychallengesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_LobbyChallengesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesMaxnotificationsubscriptiondelayseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_MaxNotificationSubscriptionDelaySeconds"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesMaxwaittimebeforenotificationsubscriptionseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_MaxWaitTimeBeforeNotificationSubscriptionSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesNumberofsuggestedchallenges() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_NumberOfSuggestedChallenges"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesPartiesv2enabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_PartiesV2Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesPostgameoverride() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_PostgameOverride"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesRankidentityoverride() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_RankIdentityOverride"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesSeasonaltooltipenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_SeasonalTooltipEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesWaittimebeforecallchallengeupdateseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_WaitTimeBeforeCallChallengeUpdateSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesChallengesWaittimebeforedarkmodeadditionalcallsseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Challenges_WaitTimeBeforeDarkModeAdditionalCallsSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigCapunlockchampionlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_CapUnlockChampionLevel"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigChampionpointqueuetypes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_ChampionPointQueueTypes"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigEndofgameenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_EndOfGameEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigGradeenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_GradeEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigLegsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_LegsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigLootchestsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_LootChestsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigMaxchampionlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_MaxChampionLevel"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigMinsummonerlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_MinSummonerLevel"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigShowgradeavailablepopup() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_ShowGradeAvailablePopup"
}
func (EventTopic) PlatformConfigV1NamespacesChampionmasteryconfigSupportedqueuetypes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionMasteryConfig_SupportedQueueTypes"
}
func (EventTopic) PlatformConfigV1NamespacesChampionselect() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionSelect"
}
func (EventTopic) PlatformConfigV1NamespacesChampionselectAllchampsavailableinaram() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionSelect_AllChampsAvailableInAram"
}
func (EventTopic) PlatformConfigV1NamespacesChampionselectAlwaysshowrewardicon() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionSelect_AlwaysShowRewardIcon"
}
func (EventTopic) PlatformConfigV1NamespacesChampionselectAutoreconnectenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionSelect_AutoReconnectEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChampionselectCollatorchampionfilterenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionSelect_CollatorChampionFilterEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChampionselectUseactionpatchv2() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionSelect_UseActionPatchV2"
}
func (EventTopic) PlatformConfigV1NamespacesChampionselectUseoptimizedbotchampionselectprocessor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionSelect_UseOptimizedBotChampionSelectProcessor"
}
func (EventTopic) PlatformConfigV1NamespacesChampionselectUseoptimizedchampselectprocessor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionSelect_UseOptimizedChampSelectProcessor"
}
func (EventTopic) PlatformConfigV1NamespacesChampionselectUseoptimizedspellselectprocessor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionSelect_UseOptimizedSpellSelectProcessor"
}
func (EventTopic) PlatformConfigV1NamespacesChampiontradeservice() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionTradeService"
}
func (EventTopic) PlatformConfigV1NamespacesChampiontradeserviceEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChampionTradeService_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesChat() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Chat"
}
func (EventTopic) PlatformConfigV1NamespacesChatdomain() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChatDomain"
}
func (EventTopic) PlatformConfigV1NamespacesChatdomainChampselectdomainname() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChatDomain_ChampSelectDomainName"
}
func (EventTopic) PlatformConfigV1NamespacesChatdomainPostgamedomainname() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ChatDomain_PostGameDomainName"
}
func (EventTopic) PlatformConfigV1NamespacesChatChathistoryenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Chat_ChatHistoryEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChatChathistorythreshold() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Chat_ChatHistoryThreshold"
}
func (EventTopic) PlatformConfigV1NamespacesChatDefaultPublicChatRooms() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Chat_Default_public_chat_rooms"
}
func (EventTopic) PlatformConfigV1NamespacesChatMaximumrostersize() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Chat_MaximumRosterSize"
}
func (EventTopic) PlatformConfigV1NamespacesChatMobileenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Chat_MobileEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesChatRenameGeneralGroupThrottle() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Chat_Rename_general_group_throttle"
}
func (EventTopic) PlatformConfigV1NamespacesChroma() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Chroma"
}
func (EventTopic) PlatformConfigV1NamespacesChromaIsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Chroma_IsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClash() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clash"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigAwardstabenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_AwardsTabEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigBracketspectateenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_BracketSpectateEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigCapacityindicatorenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_CapacityIndicatorEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigClashstartmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_ClashStartModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigEattooltipenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_EatTooltipEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigEnabledstate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_EnabledState"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigEndofgameflowenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_EndOfGameFlowEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigFindteamviewenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_FindTeamViewEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigHonorlevelrequired() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_HonorLevelRequired"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigIconconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_IconConfig"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigInvitemodaltiersenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_InviteModalTiersEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigIsplaymoderestrictionenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_IsPlaymodeRestrictionEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigIsrewardsmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_IsRewardsModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigLftenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_LftEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigLoginmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_LoginModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigPremiumticketsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_PremiumTicketsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigScoutingenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_ScoutingEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigStorepagelink() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_StorePageLink"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigThirdpartyinvitesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_ThirdPartyInvitesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigTutorialenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_TutorialEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigVisibility() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_Visibility"
}
func (EventTopic) PlatformConfigV1NamespacesClashconfigWorlds2020lootmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClashConfig_Worlds2020LootModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClashAramintromodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clash_AramIntroModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstates() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesMaxnumplayersforpracticetoolgame() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_MaxNumPlayersForPracticeToolGame"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesAdvancedtutorialenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_advancedTutorialEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesArchivedstatsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_archivedStatsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesBuddynotesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_buddyNotesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesChampiontradethroughlcds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_championTradeThroughLCDS"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesClientheartbeatrateseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_clientHeartBeatRateSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesCurrentseason() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_currentSeason"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesDisplaypromogamesplayedenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_displayPromoGamesPlayedEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesEnabledqueueidslist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_enabledQueueIdsList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesFreetoplaychampionfornewplayersidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_freeToPlayChampionForNewPlayersIdList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesFreetoplaychampionidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_freeToPlayChampionIdList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesFreetoplaychampionsfornewplayersmaxlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_freeToPlayChampionsForNewPlayersMaxLevel"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesGamemapenableddtolist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_gameMapEnabledDTOList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesGamemodetoinactivespellids() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_gameModeToInactiveSpellIds"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesInactivearamspellidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_inactiveAramSpellIdList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesInactivechampionidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_inactiveChampionIdList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesInactiveclassicspellidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_inactiveClassicSpellIdList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesInactiveodinspellidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_inactiveOdinSpellIdList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesInactivespellidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_inactiveSpellIdList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesInactivetutorialspellidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_inactiveTutorialSpellIdList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesKnowngeographicgameserverregions() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_knownGeographicGameServerRegions"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesLeagueserviceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_leagueServiceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesLeaguesdecaymessagingenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_leaguesDecayMessagingEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesLocalespecificchatroomsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_localeSpecificChatRoomsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesMasterypageonserver() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_masteryPageOnServer"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesMaxmasterypagesonserver() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_maxMasteryPagesOnServer"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesMinnumplayersforpracticegame() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_minNumPlayersForPracticeGame"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesModulargamemodeenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_modularGameModeEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesObservablecustomgamemodes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_observableCustomGameModes"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesObservablegamemodes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_observableGameModes"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesObservermodeenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_observerModeEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesPracticegameenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_practiceGameEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesPracticegametypeconfigidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_practiceGameTypeConfigIdList"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesQueuethrottledto() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_queueThrottleDTO"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesReplayserviceaddress() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_replayServiceAddress"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesReplaysystemstates() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_replaySystemStates"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesRiotdataservicedatasendprobability() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_riotDataServiceDataSendProbability"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesRuneuniqueperspellbook() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_runeUniquePerSpellBook"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesSendfeedbackeventsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_sendFeedbackEventsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesSpectatorslotlimit() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_spectatorSlotLimit"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesStorecustomerenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_storeCustomerEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesTeamserviceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_teamServiceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesTournamentsendstatsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_tournamentSendStatsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesTournamentshortcodesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_tournamentShortCodesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesTribunalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_tribunalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClientsystemstatesUnobtainablechampionskinidlist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ClientSystemStates_unobtainableChampionSkinIDList"
}
func (EventTopic) PlatformConfigV1NamespacesClubs() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clubs"
}
func (EventTopic) PlatformConfigV1NamespacesClubsClubpresencedecryptionkey() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clubs_ClubPresenceDecryptionKey"
}
func (EventTopic) PlatformConfigV1NamespacesClubsClubserviceurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clubs_ClubServiceUrl"
}
func (EventTopic) PlatformConfigV1NamespacesClubsClubsactiveicon() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clubs_ClubsActiveIcon"
}
func (EventTopic) PlatformConfigV1NamespacesClubsClubsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clubs_ClubsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClubsClubslcuenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clubs_ClubsLCUEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesClubsClubsmembericon() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clubs_ClubsMemberIcon"
}
func (EventTopic) PlatformConfigV1NamespacesClubsInvitetoclublobbyenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Clubs_InviteToClubLobbyEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesCompanions() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Companions"
}
func (EventTopic) PlatformConfigV1NamespacesCompanionsSelectorinchampselectenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Companions_SelectorInChampSelectEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesContextualeducation() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ContextualEducation"
}
func (EventTopic) PlatformConfigV1NamespacesContextualeducationurls() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ContextualEducationURLs"
}
func (EventTopic) PlatformConfigV1NamespacesContextualeducationurlsLastHit() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ContextualEducationURLs_LAST_HIT"
}
func (EventTopic) PlatformConfigV1NamespacesContextualeducationEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ContextualEducation_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesContextualeducationMaxtargetsummonerlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ContextualEducation_MaxTargetSummonerLevel"
}
func (EventTopic) PlatformConfigV1NamespacesContextualeducationTargetminionsperwave() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ContextualEducation_TargetMinionsPerWave"
}
func (EventTopic) PlatformConfigV1NamespacesCustomgame() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_CustomGame"
}
func (EventTopic) PlatformConfigV1NamespacesCustomgameBotsavailableinaram() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_CustomGame_BotsAvailableInAram"
}
func (EventTopic) PlatformConfigV1NamespacesCustomgameMinorrestrictionsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_CustomGame_MinorRestrictionsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampion() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampion"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionskins() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampionSkins"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionskinsDisabledchampionskins() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampionSkins_DisabledChampionSkins"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionskinsDisabledchromas() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampionSkins_DisabledChromas"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampions() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsAram() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ARAM"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsAramBot() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ARAM_BOT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsAramClash() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ARAM_CLASH"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsAramUnranked5x5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ARAM_UNRANKED_5x5"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsArsr() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ARSR"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsAscension() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ASCENSION"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsBilgewater() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_BILGEWATER"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsBot() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_BOT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsBot3x3() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_BOT_3x3"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsBrawl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_BRAWL"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsCherry() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_CHERRY"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsCherryUnranked() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_CHERRY_UNRANKED"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsChonccTreasureTft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_CHONCC_TREASURE_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsClash() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_CLASH"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsClassic() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_CLASSIC"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsCounterPick() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_COUNTER_PICK"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsFirstblood() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_FIRSTBLOOD"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsFirstblood1x1() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_FIRSTBLOOD_1x1"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsFirstblood2x2() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_FIRSTBLOOD_2x2"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsFiveYearAnniversaryTft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_FIVE_YEAR_ANNIVERSARY_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsGamemodex() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_GAMEMODEX"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsHexakill() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_HEXAKILL"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsKingporo() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_KINGPORO"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsKingporo5x5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_KINGPORO-5X5"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsKingPoro() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_KING_PORO"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsLny23Tft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_LNY23_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsLny24Tft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_LNY24_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsLny25Tft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_LNY25_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsNexusblitz() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_NEXUSBLITZ"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsNightmareBot() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_NIGHTMARE_BOT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsNormal() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_NORMAL"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsNormal3x3() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_NORMAL_3x3"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsNormalTft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_NORMAL_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsOdin() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ODIN"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsOdinbot() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ODINBOT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsOdinbot5x5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ODINBOT_5x5"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsOdinUnranked() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ODIN_UNRANKED"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsOdyssey() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ODYSSEY"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsOneforall() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ONEFORALL"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsOneforall5x5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ONEFORALL_5x5"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsPracticetool() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_PRACTICETOOL"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsPvePuzzleTft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_PVE_PUZZLE_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedFlexSr() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_FLEX_SR"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedFlexTt() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_FLEX_TT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedPremade3x3() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_PREMADE_3x3"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedPremade5x5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_PREMADE_5x5"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedSolo5x5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_SOLO_5x5"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedTeam3x3() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_TEAM_3x3"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedTeam5x5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_TEAM_5x5"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedTft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedTftDoubleUp() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_TFT_DOUBLE_UP"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedTftPairs() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_TFT_PAIRS"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRankedTftTurbo() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RANKED_TFT_TURBO"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsRiotscriptBot() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_RIOTSCRIPT_BOT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsSetRevival55Tft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_SET_REVIVAL_5_5_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsSetRevivalTft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_SET_REVIVAL_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsSfTft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_SF_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsSiege() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_SIEGE"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsSnowurf() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_SNOWURF"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsSr6x6() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_SR_6x6"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsStrawberry() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_STRAWBERRY"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsSwiftplay() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_SWIFTPLAY"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsTft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsTutorial() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_TUTORIAL"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsTutorialModule1() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_TUTORIAL_MODULE_1"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsTutorialModule2() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_TUTORIAL_MODULE_2"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsTutorialModule3() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_TUTORIAL_MODULE_3"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsTutorialTft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_TUTORIAL_TFT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsUltbook() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_ULTBOOK"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsUrf() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_URF"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsUrfBot() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_URF_BOT"
}
func (EventTopic) PlatformConfigV1NamespacesDisabledchampionsUrfClash() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DisabledChampions_URF_CLASH"
}
func (EventTopic) PlatformConfigV1NamespacesDiscordrpIsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DiscordRP_IsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesDockedprompt() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DockedPrompt"
}
func (EventTopic) PlatformConfigV1NamespacesDockedpromptEnablednewdockedpromptrenderer() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_DockedPrompt_EnabledNewDockedPromptRenderer"
}
func (EventTopic) PlatformConfigV1NamespacesESports() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports"
}
func (EventTopic) PlatformConfigV1NamespacesESportsCsCzEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_CS_CZ_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsDeDeEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_DE_DE_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsElGrEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EL_GR_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEnAuEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EN_AU_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEnGbEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EN_GB_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEnPhEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EN_PH_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEnPlEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EN_PL_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEnSgEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EN_SG_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEnUsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EN_US_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEsArEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_ES_AR_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEsEsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_ES_ES_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEsMxEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_ES_MX_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEsportshubdataurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EsportsHubDataURL"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEsportshubinitialized() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EsportsHubInitialized"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEsportshublongpollminutes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EsportsHubLongPollMinutes"
}
func (EventTopic) PlatformConfigV1NamespacesESportsEsportshubshortpollminutes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_EsportsHubShortPollMinutes"
}
func (EventTopic) PlatformConfigV1NamespacesESportsFrFrEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_FR_FR_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsHuHuEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_HU_HU_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsIdIdEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_ID_ID_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsItItEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_IT_IT_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsJaJpEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_JA_JP_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsKoKrEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_KO_KR_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsKillhub() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_KillHub"
}
func (EventTopic) PlatformConfigV1NamespacesESportsLandingEmbedUrl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_Landing_Embed_URL"
}
func (EventTopic) PlatformConfigV1NamespacesESportsMsMyEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_MS_MY_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsPlPlEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_PL_PL_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsPtBrEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_PT_BR_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsRoRoEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_RO_RO_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsRuRuEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_RU_RU_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsStaging() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_Staging"
}
func (EventTopic) PlatformConfigV1NamespacesESportsThThEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_TH_TH_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsTrTrEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_TR_TR_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsVnVnEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_VN_VN_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsZhCnEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_ZH_CN_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsZhMyEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_ZH_MY_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesESportsZhTwEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ESports_ZH_TW_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesEmotes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Emotes"
}
func (EventTopic) PlatformConfigV1NamespacesEmotesIsemotepanelenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Emotes_IsEmotePanelEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesEmotesIsemotetutorialmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Emotes_IsEmoteTutorialModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesEndofgamegiftsettings() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EndOfGameGiftSettings"
}
func (EventTopic) PlatformConfigV1NamespacesEndofgamegiftsettingsGiftrecipientlevelmin() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EndOfGameGiftSettings_GiftRecipientLevelMin"
}
func (EventTopic) PlatformConfigV1NamespacesEndofgamegiftsettingsGiftsenderlevelmin() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EndOfGameGiftSettings_GiftSenderLevelMin"
}
func (EventTopic) PlatformConfigV1NamespacesEndofgamegiftsettingsGiftsenderrpmax() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EndOfGameGiftSettings_GiftSenderRPMax"
}
func (EventTopic) PlatformConfigV1NamespacesEndofgamegiftsettingsRecipientgiftdailymax() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EndOfGameGiftSettings_RecipientGiftDailyMax"
}
func (EventTopic) PlatformConfigV1NamespacesEndofgamegiftsettingsSendergiftdailymax() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EndOfGameGiftSettings_SenderGiftDailyMax"
}
func (EventTopic) PlatformConfigV1NamespacesEndofgamegifting() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EndOfGameGifting"
}
func (EventTopic) PlatformConfigV1NamespacesEndofgamegiftingEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EndOfGameGifting_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesEogreporting() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EoGReporting"
}
func (EventTopic) PlatformConfigV1NamespacesEogreportingGameagnosticreportingtrinary() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_EoGReporting_GameAgnosticReportingTrinary"
}
func (EventTopic) PlatformConfigV1NamespacesEsports() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Esports"
}
func (EventTopic) PlatformConfigV1NamespacesEsportsNotificationsassetmagickurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Esports_NotificationsAssetMagickURL"
}
func (EventTopic) PlatformConfigV1NamespacesEsportsNotificationsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Esports_NotificationsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesEsportsNotificationsserviceendpoint() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Esports_NotificationsServiceEndpoint"
}
func (EventTopic) PlatformConfigV1NamespacesEsportsNotificationsstreamgroupslug() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Esports_NotificationsStreamGroupSlug"
}
func (EventTopic) PlatformConfigV1NamespacesEsportsNotificationsstreamurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Esports_NotificationsStreamURL"
}
func (EventTopic) PlatformConfigV1NamespacesEternals() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Eternals"
}
func (EventTopic) PlatformConfigV1NamespacesEternalsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Eternals_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesEternalsServiceurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Eternals_ServiceUrl"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedgame() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedGame"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedgameMetadataenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedGame_MetadataEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedmodes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedModes"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedmodesDisabledrgmbuttonenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedModes_DisabledRgmButtonEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedmodesGoldenspatulaclubdisabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedModes_GoldenSpatulaClubDisabled"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedmodesMaxnotificationsavedelayminutes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedModes_MaxNotificationSaveDelayMinutes"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedmodesNotificationsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedModes_NotificationsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedmodesQueuetogglenotificationminutesthreshold() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedModes_QueueToggleNotificationMinutesThreshold"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedmodesQueuesdelayedrefreshmaxtimeout() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedModes_QueuesDelayedRefreshMaxTimeout"
}
func (EventTopic) PlatformConfigV1NamespacesFeaturedmodesQueuesdelayedrefreshmintimeout() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FeaturedModes_QueuesDelayedRefreshMinTimeout"
}
func (EventTopic) PlatformConfigV1NamespacesFriendrecommendations() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_FriendRecommendations"
}
func (EventTopic) PlatformConfigV1NamespacesGameclientcmdlinetoggles() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameClientCmdLineToggles"
}
func (EventTopic) PlatformConfigV1NamespacesGameclientcmdlinetogglesUsedx11() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameClientCmdLineToggles_UseDX11"
}
func (EventTopic) PlatformConfigV1NamespacesGameinvites() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameInvites"
}
func (EventTopic) PlatformConfigV1NamespacesGameinvitesGameinviteserviceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameInvites_GameInviteServiceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesGameinvitesInvitebulkmaxsize() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameInvites_InviteBulkMaxSize"
}
func (EventTopic) PlatformConfigV1NamespacesGameinvitesLobbycreationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameInvites_LobbyCreationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesGameinvitesServiceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameInvites_ServiceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesGametimersync() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameTimerSync"
}
func (EventTopic) PlatformConfigV1NamespacesGametimersyncEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameTimerSync_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesGametimersyncPercentoftotaltimertosyncat() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GameTimerSync_PercentOfTotalTimerToSyncAt"
}
func (EventTopic) PlatformConfigV1NamespacesGameflow() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Gameflow"
}
func (EventTopic) PlatformConfigV1NamespacesGameflowForcegamelocaleasenglish() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Gameflow_ForceGameLocaleAsEnglish"
}
func (EventTopic) PlatformConfigV1NamespacesGameflowShouldsendriotclientheartbeat() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Gameflow_ShouldSendRiotClientHeartBeat"
}
func (EventTopic) PlatformConfigV1NamespacesGeoinfo() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GeoInfo"
}
func (EventTopic) PlatformConfigV1NamespacesGuestslots() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_GuestSlots"
}
func (EventTopic) PlatformConfigV1NamespacesHighlights() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Highlights"
}
func (EventTopic) PlatformConfigV1NamespacesHighlightsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Highlights_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesHonor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Honor"
}
func (EventTopic) PlatformConfigV1NamespacesHonorDayonemodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Honor_DayOneModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesHonorEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Honor_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesHonorHonor2018enabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Honor_Honor2018Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesHonorSecondstovote() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Honor_SecondsToVote"
}
func (EventTopic) PlatformConfigV1NamespacesIconconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_IconConfig"
}
func (EventTopic) PlatformConfigV1NamespacesIconconfigEnabledstate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_IconConfig_EnabledState"
}
func (EventTopic) PlatformConfigV1NamespacesInventory() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Inventory"
}
func (EventTopic) PlatformConfigV1NamespacesInventoryBaseserviceurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Inventory_BaseServiceUrl"
}
func (EventTopic) PlatformConfigV1NamespacesInventoryEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Inventory_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesItemsets() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ItemSets"
}
func (EventTopic) PlatformConfigV1NamespacesItemsetsEditorenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ItemSets_EditorEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesItemsetsMaxitemsets() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ItemSets_MaxItemSets"
}
func (EventTopic) PlatformConfigV1NamespacesItemsetsSenditemsetstogame() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ItemSets_SendItemSetsToGame"
}
func (EventTopic) PlatformConfigV1NamespacesKarma() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Karma"
}
func (EventTopic) PlatformConfigV1NamespacesKarmaWorlds2017votingenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Karma_Worlds2017VotingEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesKickout() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Kickout"
}
func (EventTopic) PlatformConfigV1NamespacesKrplaytimereminder() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_KrPlaytimeReminder"
}
func (EventTopic) PlatformConfigV1NamespacesLcu() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCU"
}
func (EventTopic) PlatformConfigV1NamespacesLcuacs() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUACS"
}
func (EventTopic) PlatformConfigV1NamespacesLcuacsEndpoint() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUACS_Endpoint"
}
func (EventTopic) PlatformConfigV1NamespacesLcuacsUseidtokens() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUACS_UseIdTokens"
}
func (EventTopic) PlatformConfigV1NamespacesLcucollections() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUCollections"
}
func (EventTopic) PlatformConfigV1NamespacesLcucollectionsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUCollections_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcucollectionsLcuaugmentsvisible() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUCollections_LCUAugmentsVisible"
}
func (EventTopic) PlatformConfigV1NamespacesLcucollectionsLcucollectiblesfinishersenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUCollections_LCUCollectiblesFinishersEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcucollectionsLcumasteriesvisible() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUCollections_LCUMasteriesVisible"
}
func (EventTopic) PlatformConfigV1NamespacesLcucollectionsLcuperksvisible() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUCollections_LCUPerksVisible"
}
func (EventTopic) PlatformConfigV1NamespacesLcucollectionsLcurunesvisible() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUCollections_LCURunesVisible"
}
func (EventTopic) PlatformConfigV1NamespacesLcucollectionsLcuskinsviewerenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUCollections_LCUSkinsViewerEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcustore() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreCanqueryinactiveitems() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_CanQueryInactiveItems"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreDisablecaprms() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_DisableCapRMS"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreEnabledropratesinpurchasemodal() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_EnableDropRatesInPurchaseModal"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreEnablefetchoffers() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_EnableFetchOffers"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreLookupmissingcatalogitemkeys() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_LookupMissingCatalogItemKeys"
}
func (EventTopic) PlatformConfigV1NamespacesLcustorePlayergiftingnotificationsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_PlayerGiftingNotificationsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreRecommendationsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_RecommendationsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreSinglepageapplicationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_SinglePageApplicationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreUsegamedataassets() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_UseGameDataAssets"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreUserms() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_UseRMS"
}
func (EventTopic) PlatformConfigV1NamespacesLcustoreUsersoaccesstoken() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCUStore_UseRsoAccessToken"
}
func (EventTopic) PlatformConfigV1NamespacesLcuAirclientalphainviteenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCU_AirClientAlphaInviteEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuAirclientalphainviteurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LCU_AirClientAlphaInviteURL"
}
func (EventTopic) PlatformConfigV1NamespacesLcualphashutdown() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuAlphaShutdown"
}
func (EventTopic) PlatformConfigV1NamespacesLcualphashutdownCountdown() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuAlphaShutdown_Countdown"
}
func (EventTopic) PlatformConfigV1NamespacesLcualphashutdownEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuAlphaShutdown_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcubuddyspectate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuBuddySpectate"
}
func (EventTopic) PlatformConfigV1NamespacesLcubuddyspectateEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuBuddySpectate_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampiondetails() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionDetails"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampiondetailsAbilitiessectionenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionDetails_AbilitiesSectionEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampiondetailsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionDetails_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampiondetailsPawenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionDetails_PawEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampiondetailsSkinssectionenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionDetails_SkinsSectionEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselect() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectAlliedskindisplayenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_AlliedSkinDisplayEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectChampselectchangetooltipdurationmillis() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_ChampSelectChangeTooltipDurationMillis"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectChampselectchangetooltipkey() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_ChampSelectChangeTooltipKey"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectChampselectmutingenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_ChampSelectMutingEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectChampselectreportv2enabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_ChampSelectReportV2Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectChamptradingtooltipenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_ChampTradingTooltipEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectDisableautosmiteassignment() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_DisableAutoSmiteAssignment"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectEnablefavorites() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_EnableFavorites"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectEnablepositionfilters() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_EnablePositionFilters"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectIsdisconnectnotificationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_IsDisconnectNotificationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectPickorderswappingtooltipenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_PickOrderSwappingTooltipEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectPositionassignmentanimationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_PositionAssignmentAnimationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectRandomchampionenabledgamequeues() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_RandomChampionEnabledGameQueues"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectRandomchampionratelimitinterval() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_RandomChampionRateLimitInterval"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectRandomchampionratelimitmaxactions() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_RandomChampionRateLimitMaxActions"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectReportingenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_ReportingEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectShowchampselectchangetooltip() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_ShowChampSelectChangeTooltip"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectShowemotebutton() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_ShowEmoteButton"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectSkinpurchaseenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_SkinPurchaseEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectSkinpurchasenotificationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_SkinPurchaseNotificationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuchampionselectSkinpurchasetime() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuChampionSelect_SkinPurchaseTime"
}
func (EventTopic) PlatformConfigV1NamespacesLcucontenttargeting() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuContentTargeting"
}
func (EventTopic) PlatformConfigV1NamespacesLcudisambiguation() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuDisambiguation"
}
func (EventTopic) PlatformConfigV1NamespacesLcuemailverification() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuEmailVerification"
}
func (EventTopic) PlatformConfigV1NamespacesLcuemailverificationEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuEmailVerification_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuemailverificationIsoptional() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuEmailVerification_IsOptional"
}
func (EventTopic) PlatformConfigV1NamespacesLcuemailverificationMandatoryat() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuEmailVerification_MandatoryAt"
}
func (EventTopic) PlatformConfigV1NamespacesLcuemailverificationMaxoptionallevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuEmailVerification_MaxOptionalLevel"
}
func (EventTopic) PlatformConfigV1NamespacesLcuemailverificationMinimumsummonerlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuEmailVerification_MinimumSummonerLevel"
}
func (EventTopic) PlatformConfigV1NamespacesLcuemailverificationRequiredlogins() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuEmailVerification_RequiredLogins"
}
func (EventTopic) PlatformConfigV1NamespacesLcugamesettings() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuGameSettings"
}
func (EventTopic) PlatformConfigV1NamespacesLcugamesettingsGameplayenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuGameSettings_GameplayEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcugamesettingsHotkeysenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuGameSettings_HotkeysEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcugamesettingsInterfaceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuGameSettings_InterfaceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcugamesettingsSoundenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuGameSettings_SoundEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuhome() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuHome"
}
func (EventTopic) PlatformConfigV1NamespacesLcuhomeLandingpagestimeout() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuHome_LandingPagesTimeout"
}
func (EventTopic) PlatformConfigV1NamespacesLcuhomeReloadenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuHome_ReloadEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuhomeReloadpollerinterval() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuHome_ReloadPollerInterval"
}
func (EventTopic) PlatformConfigV1NamespacesLcuhomeReloadstaleinterval() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuHome_ReloadStaleInterval"
}
func (EventTopic) PlatformConfigV1NamespacesLcuhomeRequireitemloaded() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuHome_RequireItemLoaded"
}
func (EventTopic) PlatformConfigV1NamespacesLcuhovercard() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuHovercard"
}
func (EventTopic) PlatformConfigV1NamespacesLcuhovercardDisabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuHovercard_Disabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcuhovercardRoleinfoenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuHovercard_RoleInfoEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculeaguespectate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLeagueSpectate"
}
func (EventTopic) PlatformConfigV1NamespacesLculeaguespectateEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLeagueSpectate_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculevelup() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLevelUp"
}
func (EventTopic) PlatformConfigV1NamespacesLculevelupModalnotificationdisabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLevelUp_ModalNotificationDisabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculobby() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby"
}
func (EventTopic) PlatformConfigV1NamespacesLculobbyAutograntinviteenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby_AutoGrantInviteEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculobbyGameinvitesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby_GameInvitesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculobbyObservermodeenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby_ObserverModeEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculobbyPracticegameenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby_PracticeGameEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculobbyPracticegamelistenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby_PracticeGameListEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculobbyQueueeligibilitygatekeeperenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby_QueueEligibilityGateKeeperEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculobbyQueueeligibilityv2enabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby_QueueEligibilityV2Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculobbySendinventorytokenmetricsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby_SendInventoryTokenMetricsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculobbyUseinventorytoken() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLobby_UseInventoryToken"
}
func (EventTopic) PlatformConfigV1NamespacesLculoyalty() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLoyalty"
}
func (EventTopic) PlatformConfigV1NamespacesLculoyaltyLeagueunlockedenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLoyalty_LeagueUnlockedEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLculoyaltyLolcafeenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuLoyalty_LolcafeEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcunpe() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuNpe"
}
func (EventTopic) PlatformConfigV1NamespacesLcunpeHardmaxsummonerlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuNpe_HardMaxSummonerLevel"
}
func (EventTopic) PlatformConfigV1NamespacesLcunpeMaxsummonerlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuNpe_MaxSummonerLevel"
}
func (EventTopic) PlatformConfigV1NamespacesLcunpeRewardschallengesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuNpe_RewardsChallengesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcunpeRewardsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuNpe_RewardsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcunpeRewardsloginenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuNpe_RewardsLoginEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcupayments() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPayments"
}
func (EventTopic) PlatformConfigV1NamespacesLcupaymentsBypassaccountids() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPayments_BypassAccountIds"
}
func (EventTopic) PlatformConfigV1NamespacesLcupaymentsHost() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPayments_Host"
}
func (EventTopic) PlatformConfigV1NamespacesLcupaymentsPmcedgehost() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPayments_PmcEdgeHost"
}
func (EventTopic) PlatformConfigV1NamespacesLcupaymentsPmcsessionsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPayments_PmcSessionsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcupaymentsRiotpayenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPayments_RiotPayEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcupaymentsRiotpaythrottle() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPayments_RiotPayThrottle"
}
func (EventTopic) PlatformConfigV1NamespacesLcuprofiles() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuProfiles"
}
func (EventTopic) PlatformConfigV1NamespacesLcuprofilesEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuProfiles_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcupurchasewidget() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPurchaseWidget"
}
func (EventTopic) PlatformConfigV1NamespacesLcupurchasewidgetBaseurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPurchaseWidget_BaseUrl"
}
func (EventTopic) PlatformConfigV1NamespacesLcupurchasewidgetCapordersurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPurchaseWidget_CapOrdersUrl"
}
func (EventTopic) PlatformConfigV1NamespacesLcupurchasewidgetEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPurchaseWidget_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcupurchasewidgetNonrefundabledisclaimerenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPurchaseWidget_NonRefundableDisclaimerEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcupurchasewidgetPurchasedisclaimerenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuPurchaseWidget_PurchaseDisclaimerEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcureadycheckupdate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuReadyCheckUpdate"
}
func (EventTopic) PlatformConfigV1NamespacesLcureadycheckupdateEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuReadyCheckUpdate_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcurevivals() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuRevivals"
}
func (EventTopic) PlatformConfigV1NamespacesLcurevivalsFoundationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuRevivals_FoundationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusentryjserrors() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSentryJSErrors"
}
func (EventTopic) PlatformConfigV1NamespacesLcusentryjserrorsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSentryJSErrors_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusentryjserrorsSamplerate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSentryJSErrors_SampleRate"
}
func (EventTopic) PlatformConfigV1NamespacesLcusentryjserrorsSentrydsn() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSentryJSErrors_SentryDSN"
}
func (EventTopic) PlatformConfigV1NamespacesLcusettings() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSettings"
}
func (EventTopic) PlatformConfigV1NamespacesLcusettingsFullrepairenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSettings_FullRepairEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocial() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialAutolinkenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_AutoLinkEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialChatwindowresizeenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_ChatWindowResizeEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialClearchathistoryenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_ClearChatHistoryEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialDefaultgamequeues() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_DefaultGameQueues"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialEnabledgamequeues() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_EnabledGameQueues"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialFriendrequesttoastsdisabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_FriendRequestToastsDisabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialFriendslistgiftingenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_FriendsListGiftingEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialLcusupportedgamequeues() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_LcuSupportedGameQueues"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialMoreunreadsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_MoreUnreadsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialNewchatbuttonenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_NewChatButtonEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialRecentlyplayeddisabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_RecentlyPlayedDisabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialSlashmeenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_SlashMeEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialSortconversationsbytimeenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_SortConversationsByTimeEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialStatusesdisabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_StatusesDisabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialVirtualizedmessagesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_VirtualizedMessagesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusocialVirtualizedrosterenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSocial_VirtualizedRosterEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcusummonericonpicker() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSummonerIconPicker"
}
func (EventTopic) PlatformConfigV1NamespacesLcusummonericonpickerEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuSummonerIconPicker_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcutft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTft"
}
func (EventTopic) PlatformConfigV1NamespacesLcutftBattlepasshubenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTft_BattlepassHubEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcutftEligibilityinventorytypes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTft_EligibilityInventoryTypes"
}
func (EventTopic) PlatformConfigV1NamespacesLcutftHubenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTft_HubEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcutftMatchhistoryenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTft_MatchHistoryEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcutftOrbenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTft_OrbEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcutftPlaybuttonenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTft_PlayButtonEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcutftSeriesinternalname() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTft_SeriesInternalName"
}
func (EventTopic) PlatformConfigV1NamespacesLcutftTfthomeenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTft_TftHomeEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorial() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorialCarouselchampids() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial_CarouselChampIds"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorialEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorialGamemodeselectenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial_GameModeSelectEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorialIntroabtestpercentage() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial_IntroABTestPercentage"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorialNewplayerexperienceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial_NewPlayerExperienceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorialPracticetoolenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial_PracticeToolEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorialSkiptutorialpathafterlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial_SkipTutorialPathAfterLevel"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorialStatstimeout() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial_StatsTimeout"
}
func (EventTopic) PlatformConfigV1NamespacesLcututorialTutorialsummonericon() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuTutorial_TutorialSummonerIcon"
}
func (EventTopic) PlatformConfigV1NamespacesLcuuikit() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuUiKit"
}
func (EventTopic) PlatformConfigV1NamespacesLcuuikitCelebrationmodalsdisabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LcuUiKit_CelebrationModalsDisabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigApexdemotionnotificationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_ApexDemotionNotificationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigChallengerladdersenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_ChallengerLaddersEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigConfigrefreshintervalseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_ConfigRefreshIntervalSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigCurrentsplit() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_CurrentSplit"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigCurrentyear() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_CurrentYear"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigDefaultjwttimetoliveseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_DefaultJwtTimeToLiveSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigEosnotificationsettingsname() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_EosNotificationSettingsName"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigEosnotificationsettingsschemaver() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_EosNotificationSettingsSchemaVer"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigEosnotificationsconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_EosNotificationsConfig"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigEosnotificationsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_EosNotificationsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigEosrewardgroupsconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_EosRewardGroupsConfig"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigEosrewardsconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_EosRewardsConfig"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigFlexrestrictionmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_FlexRestrictionModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigIsglobalnotificationsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_IsGlobalNotificationsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigIspreseason() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_IsPreseason"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigIsseasonmemorialmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_IsSeasonMemorialModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigIssplitstartmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_IsSplitStartModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigJwtenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_JWTEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigLeagueserviceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_LeagueServiceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigMastertierenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_MasterTierEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigPreseasonname() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_PreseasonName"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigPromohelperenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_PromoHelperEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigPromotionvignettev2enabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_PromotionVignetteV2Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigRanked2017enabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_Ranked2017Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigRanked2019enabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_Ranked2019Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigRankedreferencemodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_RankedReferenceModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigRankedrewardconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_RankedRewardConfig"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigSeasonmemorialmodalminhonorlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_SeasonMemorialModalMinHonorLevel"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigSeasonmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_SeasonModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigSeasonname() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_SeasonName"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigSendsignedrankedoverview() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_SendSignedRankedOverview"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueconfigTftseasonnamelockey() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueConfig_TftSeasonNameLocKey"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclient() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClient"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservices() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesGsmv2() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_GSMv2"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesGameagnosticmatchhistory() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_GameAgnosticMatchHistory"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesLeagues() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_Leagues"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesMarketingpreferences() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_MarketingPreferences"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesMissions() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_Missions"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesParties() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_Parties"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesPlayerbehavior() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_PlayerBehavior"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesPublishingcontent() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_PublishingContent"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesSummoner() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_Summoner"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesTastes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_Tastes"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientenabledservicesTeambuilder() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientEnabledServices_Teambuilder"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientleaguesessionservicesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientLeagueSessionServicesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientservicetrafficbalancerate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClientServiceTrafficBalanceRate"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClient_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeagueedgeclientLeaguesessionenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueEdgeClient_LeagueSessionEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeaguesessionFailurerefreshtimeoutseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueSession_FailureRefreshTimeoutSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesLeaguesessionRefreshtokenoverride() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueSession_RefreshTokenOverride"
}
func (EventTopic) PlatformConfigV1NamespacesLeaguesessionUsesessionrefreshv2() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeagueSession_UseSessionRefreshV2"
}
func (EventTopic) PlatformConfigV1NamespacesLeaverbuster() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeaverBuster"
}
func (EventTopic) PlatformConfigV1NamespacesLeaverbusterIslbsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeaverBuster_IsLbsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLeaverbusterIslockoutmodalenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LeaverBuster_IsLockoutModalEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLoadouts() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Loadouts"
}
func (EventTopic) PlatformConfigV1NamespacesLoadoutsEnablestarshardsservices() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Loadouts_EnableStarShardsServices"
}
func (EventTopic) PlatformConfigV1NamespacesLoadoutsEnablestarshardsupgradeflow() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Loadouts_EnableStarShardsUpgradeFlow"
}
func (EventTopic) PlatformConfigV1NamespacesLoadoutsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Loadouts_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLoadoutsInventoryserviceurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Loadouts_InventoryServiceUrl"
}
func (EventTopic) PlatformConfigV1NamespacesLoadoutsLoadoutsserviceurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Loadouts_LoadoutsServiceUrl"
}
func (EventTopic) PlatformConfigV1NamespacesLoadoutsNewloadoutpickerenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Loadouts_NewLoadoutPickerEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLoadoutsUsev4loadoutflow() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Loadouts_UseV4LoadoutFlow"
}
func (EventTopic) PlatformConfigV1NamespacesLoadoutsValidinventorytypes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Loadouts_ValidInventoryTypes"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacket() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketBanneduntildatemillis() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_bannedUntilDateMillis"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketBroadcastnotification() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_broadcastNotification"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketCoopvsaiminuteslefttoday() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_coOpVsAiMinutesLeftToday"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketCoopvsaimsecsuntilreset() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_coOpVsAiMsecsUntilReset"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketCompetitiveregion() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_competitiveRegion"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketCustomminuteslefttoday() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_customMinutesLeftToday"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketCustommsecsuntilreset() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_customMsecsUntilReset"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketDisplayprimereformcard() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_displayPrimeReformCard"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketEmailstatus() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_emailStatus"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketGametypeconfigs() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_gameTypeConfigs"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketInghostgame() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_inGhostGame"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketLanguages() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_languages"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketMatchmakingenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_matchMakingEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketMaxpracticegamesize() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_maxPracticeGameSize"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketMinor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_minor"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketMinorshutdownenforced() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_minorShutdownEnforced"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketMinutesuntilmidnight() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_minutesUntilMidnight"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketMinutesuntilshutdown() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_minutesUntilShutdown"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketMinutesuntilshutdownenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_minutesUntilShutdownEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketPlatformgamelifecycle() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_platformGameLifecycle"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketPlatformid() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_platformId"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketPlayerstatsummaries() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_playerStatSummaries"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketRestrictedchatgamesremaining() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_restrictedChatGamesRemaining"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketRestrictedgamesremainingforranked() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_restrictedGamesRemainingForRanked"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketShowemailverificationpopup() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_showEmailVerificationPopup"
}
func (EventTopic) PlatformConfigV1NamespacesLogindatapacketSimplemessages() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LoginDataPacket_simpleMessages"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigChestbundlediscount1() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_ChestBundleDiscount1"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigChestbundlediscount2() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_ChestBundleDiscount2"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigChestbundlediscount3() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_ChestBundleDiscount3"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigChestbundlediscount4() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_ChestBundleDiscount4"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigChestbundlediscount5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_ChestBundleDiscount5"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigEnablemythicessencedisplay() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_EnableMythicEssenceDisplay"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigEventchestbundleid1() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_EventChestBundleId1"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigEventchestbundleid2() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_EventChestBundleId2"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigEventchestbundleid3() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_EventChestBundleId3"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigEventchestbundleid4() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_EventChestBundleId4"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigEventchestbundleid5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_EventChestBundleId5"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigEventchestsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_EventChestsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigInitializationgoalflags() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_InitializationGoalFlags"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigLedgeaccessenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_LEdgeAccessEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigLedgeaccesspercentage() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_LEdgeAccessPercentage"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigLcuenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_LcuEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigLootmilestonesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_LootMilestonesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigLootoddsqueryevaluationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_LootOddsQueryEvaluationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigMaterial112learnmoreurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_MATERIAL_112LearnMoreURL"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigMinsummonerlevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_MinSummonerLevel"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigNewplayerchestenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_NewPlayerChestEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigNorefundconfirmationenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_NoRefundConfirmationEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigPurchasechestsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_PurchaseChestsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigVisible() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_Visible"
}
func (EventTopic) PlatformConfigV1NamespacesLootconfigWorldstokensenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_LootConfig_WorldsTokensEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesMasteries() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Masteries"
}
func (EventTopic) PlatformConfigV1NamespacesMasteriesShowpointsresetmessage() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Masteries_ShowPointsResetMessage"
}
func (EventTopic) PlatformConfigV1NamespacesMissions() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions"
}
func (EventTopic) PlatformConfigV1NamespacesMissionsEligibilityinventorytypes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions_EligibilityInventoryTypes"
}
func (EventTopic) PlatformConfigV1NamespacesMissionsMissionscompressed() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions_MissionsCompressed"
}
func (EventTopic) PlatformConfigV1NamespacesMissionsMissionscooldownpollingtime() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions_MissionsCooldownPollingTime"
}
func (EventTopic) PlatformConfigV1NamespacesMissionsMissionsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions_MissionsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesMissionsMissionsfrontendenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions_MissionsFrontEndEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesMissionsMissionspollingtime() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions_MissionsPollingTime"
}
func (EventTopic) PlatformConfigV1NamespacesMissionsMissionstabtrackerenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions_MissionsTabTrackerEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesMissionsMissionsusev4api() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions_MissionsUseV4Api"
}
func (EventTopic) PlatformConfigV1NamespacesMissionsSendsimpleinventorytokens() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Missions_SendSimpleInventoryTokens"
}
func (EventTopic) PlatformConfigV1NamespacesMutators() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Mutators"
}
func (EventTopic) PlatformConfigV1NamespacesMutatorsBasictutorialmutators() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Mutators_BasicTutorialMutators"
}
func (EventTopic) PlatformConfigV1NamespacesMutatorsBattletrainingmutators() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Mutators_BattleTrainingMutators"
}
func (EventTopic) PlatformConfigV1NamespacesMutatorsCustomgamemutators() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Mutators_CustomGameMutators"
}
func (EventTopic) PlatformConfigV1NamespacesMutatorsEnabledassetmutators() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Mutators_EnabledAssetMutators"
}
func (EventTopic) PlatformConfigV1NamespacesMutatorsEnabledcustommodes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Mutators_EnabledCustomModes"
}
func (EventTopic) PlatformConfigV1NamespacesMutatorsEnabledmodes() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Mutators_EnabledModes"
}
func (EventTopic) PlatformConfigV1NamespacesMutatorsEnabledmutators() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Mutators_EnabledMutators"
}
func (EventTopic) PlatformConfigV1NamespacesNavigation() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Navigation"
}
func (EventTopic) PlatformConfigV1NamespacesNavigationNavbardisplaymode() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Navigation_NavBarDisplayMode"
}
func (EventTopic) PlatformConfigV1NamespacesNavigationUseenhancedmenu() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Navigation_UseEnhancedMenu"
}
func (EventTopic) PlatformConfigV1NamespacesNewmatchhistory() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewMatchHistory"
}
func (EventTopic) PlatformConfigV1NamespacesNewmatchhistoryEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewMatchHistory_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesNewmatchhistoryMatchhistoryweburl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewMatchHistory_MatchHistoryWebURL"
}
func (EventTopic) PlatformConfigV1NamespacesNewmatchhistoryPostipxptolegsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewMatchHistory_PostIPXPToLegSEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesNewmatchhistoryRecentlyplayednumgames() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewMatchHistory_RecentlyPlayedNumGames"
}
func (EventTopic) PlatformConfigV1NamespacesNewmatchhistorySharingenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewMatchHistory_SharingEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesNewmatchhistoryTftmatchhistoryenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewMatchHistory_TftMatchHistoryEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesNewplayerintro() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewPlayerIntro"
}
func (EventTopic) PlatformConfigV1NamespacesNewplayerintroIntrourl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewPlayerIntro_IntroUrl"
}
func (EventTopic) PlatformConfigV1NamespacesNewplayerintroNewsummonericonids() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewPlayerIntro_NewSummonerIconIds"
}
func (EventTopic) PlatformConfigV1NamespacesNewplayerrouter() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewPlayerRouter"
}
func (EventTopic) PlatformConfigV1NamespacesNewplayerrouterAbdisablingoftutorial() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewPlayerRouter_ABDisablingOfTutorial"
}
func (EventTopic) PlatformConfigV1NamespacesNewplayerrouterQueueid() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_NewPlayerRouter_QueueID"
}
func (EventTopic) PlatformConfigV1NamespacesParties() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesEnablelobbycreation() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_EnableLobbyCreation"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesEnabledforteambuilderqueues() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_EnabledForTeamBuilderQueues"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesGameflowregistrationstatusrequired() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_GameflowRegistrationStatusRequired"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesGameflowsamplingprobability() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_GameflowSamplingProbability"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesLoginregistrationtimeout() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_LoginRegistrationTimeout"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesNotificationdelaysamplingpercentage() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_NotificationDelaySamplingPercentage"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesOpenpartyenable() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_OpenPartyEnable"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesPremadeeligibilityfrompartiesenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_PremadeEligibilityFromPartiesEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesPremadeeligibilityqueuesdelayedrefreshtimeoutseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_PremadeEligibilityQueuesDelayedRefreshTimeoutSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesRegistrationconfigurationchangedtimeout() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_RegistrationConfigurationChangedTimeout"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesRegistrationcredentialsrequired() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_RegistrationCredentialsRequired"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesRegistrationretryinterval() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_RegistrationRetryInterval"
}
func (EventTopic) PlatformConfigV1NamespacesPartiesServiceproxysamplingpercentage() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Parties_ServiceProxySamplingPercentage"
}
func (EventTopic) PlatformConfigV1NamespacesPartyrewards() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PartyRewards"
}
func (EventTopic) PlatformConfigV1NamespacesPartyrewardsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PartyRewards_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesPartyrewardsGameflowvisualupdate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PartyRewards_GameFlowVisualUpdate"
}
func (EventTopic) PlatformConfigV1NamespacesPerks() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Perks"
}
func (EventTopic) PlatformConfigV1NamespacesPerksMinsummonerlevelunlockcustompages() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Perks_MinSummonerLevelUnlockCustomPages"
}
func (EventTopic) PlatformConfigV1NamespacesPerksPerksenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Perks_PerksEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffers() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersBaseserviceurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_BaseServiceURL"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersHubenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_HubEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersPort() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_Port"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersPromotionendtime() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_PromotionEndTime"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersPromotionname() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_PromotionName"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersPromotionstarttime() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_PromotionStartTime"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersProtocol() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_Protocol"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersServicepath() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_ServicePath"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersShownavbutton() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_ShowNavButton"
}
func (EventTopic) PlatformConfigV1NamespacesPersonalizedoffersThemedbackground() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonalizedOffers_ThemedBackground"
}
func (EventTopic) PlatformConfigV1NamespacesPersonlizedoffers() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonlizedOffers"
}
func (EventTopic) PlatformConfigV1NamespacesPersonlizedoffersBaseserviceurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonlizedOffers_BaseServiceURL"
}
func (EventTopic) PlatformConfigV1NamespacesPersonlizedoffersHubenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonlizedOffers_HubEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesPersonlizedoffersPort() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonlizedOffers_Port"
}
func (EventTopic) PlatformConfigV1NamespacesPersonlizedoffersProtocol() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonlizedOffers_Protocol"
}
func (EventTopic) PlatformConfigV1NamespacesPersonlizedoffersServiceconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonlizedOffers_ServiceConfig"
}
func (EventTopic) PlatformConfigV1NamespacesPersonlizedoffersServicepath() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PersonlizedOffers_ServicePath"
}
func (EventTopic) PlatformConfigV1NamespacesPlatformshutdown() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlatformShutdown"
}
func (EventTopic) PlatformConfigV1NamespacesPlatformshutdownEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlatformShutdown_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerbehavior() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerBehavior"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerbehaviorCodeofconductenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerBehavior_CodeOfConductEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerbehaviorReformcardv2enabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerBehavior_ReformCardV2Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerfeedbacktool() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerFeedbackTool"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerfeedbacktoolBackendurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerFeedbackTool_BackendUrl"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerfeedbacktoolEnablehometrigger() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerFeedbackTool_EnableHomeTrigger"
}
func (EventTopic) PlatformConfigV1NamespacesPlayernotification() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerNotification"
}
func (EventTopic) PlatformConfigV1NamespacesPlayernotificationEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerNotification_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerpreferences() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerPreferences"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerpreferencesEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerPreferences_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerpreferencesEnforcehttps() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerPreferences_EnforceHttps"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerpreferencesServiceendpoint() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerPreferences_ServiceEndpoint"
}
func (EventTopic) PlatformConfigV1NamespacesPlayerpreferencesVersion() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PlayerPreferences_Version"
}
func (EventTopic) PlatformConfigV1NamespacesPostgame() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Postgame"
}
func (EventTopic) PlatformConfigV1NamespacesPostgameShowpositiondetectionenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Postgame_ShowPositionDetectionEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercard() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardAcslookup() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_ACSLookup"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardChampionmasterylookup() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_ChampionMasteryLookup"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardIsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_IsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardIsenabledforbuddypanelview() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_IsEnabledForBuddyPanelView"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardIsenabledforchatfriendslist() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_IsEnabledForChatFriendsList"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardIsenabledforchatgroupchatparticipants() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_IsEnabledForChatGroupChatParticipants"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardIsenabledforclubchatparticipants() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_IsEnabledForClubChatParticipants"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardIsenabledforcustomgamelobby() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_IsEnabledForCustomGameLobby"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardIsenabledforgamelobbysuggestedplayers() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_IsEnabledForGameLobbySuggestedPlayers"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardIsenabledforsummonerquickview() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_IsEnabledForSummonerQuickView"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardIsenabledforteambuildersuggestedplayers() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_IsEnabledForTeamBuilderSuggestedPlayers"
}
func (EventTopic) PlatformConfigV1NamespacesProfilehovercardLeaguelookup() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ProfileHoverCard_LeagueLookup"
}
func (EventTopic) PlatformConfigV1NamespacesProfiles() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Profiles"
}
func (EventTopic) PlatformConfigV1NamespacesProfilesEosiconenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Profiles_EosIconEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesProfilesEostooltipshowonce() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Profiles_EosTooltipShowOnce"
}
func (EventTopic) PlatformConfigV1NamespacesProfilesRankqueueoncount() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Profiles_RankQueueOnCount"
}
func (EventTopic) PlatformConfigV1NamespacesProfilesSkinspickerenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Profiles_SkinsPickerEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesPublishingcontent() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PublishingContent"
}
func (EventTopic) PlatformConfigV1NamespacesPublishingcontentEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PublishingContent_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesPublishingcontentLocalepreferenceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PublishingContent_LocalePreferenceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesPublishingcontentLocalepreferenceoptions() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PublishingContent_LocalePreferenceOptions"
}
func (EventTopic) PlatformConfigV1NamespacesPublishingcontentTfthubcardsurl() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_PublishingContent_TftHubCardsUrl"
}
func (EventTopic) PlatformConfigV1NamespacesQueueimages() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueImages"
}
func (EventTopic) PlatformConfigV1NamespacesQueueimagesOverridequeueimage() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueImages_OverrideQueueImage"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestriction() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionAllowablepremadesizesforqueueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_AllowablePremadeSizesForQueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionKarmaenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_KarmaEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtieranddivisionforpremadesize2queueid1100() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndDivisionForPremadeSize2QueueId1100"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtieranddivisionforpremadesize2queueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndDivisionForPremadeSize2QueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize1() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize1"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize1queueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize1QueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize1queueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize1QueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize1queueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize1QueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize1queueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize1QueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize1queueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize1QueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize2() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize2"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize2queueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize2QueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize2queueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize2QueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize2queueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize2QueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize2queueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize2QueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize2queueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize2QueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize3() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize3"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize3queueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize3QueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize3queueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize3QueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize3queueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize3QueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize3queueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize3QueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize3queueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize3QueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize4() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize4"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize4queueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize4QueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize4queueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize4QueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize4queueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize4QueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize4queueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize4QueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize4queueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize4QueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize5() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize5"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize5queueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize5QueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize5queueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize5QueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize5queueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize5QueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize5queueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize5QueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionMaxtierandrankforpremadesize5queueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_MaxTierAndRankForPremadeSize5QueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionQueuesrequiringtwentychampions() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_QueuesRequiringTwentyChampions"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedduoqueuedefaultunseededtier() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedDuoQueueDefaultUnseededTier"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedduoqueuedefaultunseededtierqueueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedDuoQueueDefaultUnseededTierQueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedduoqueuedefaultunseededtierqueueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedDuoQueueDefaultUnseededTierQueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedduoqueuedefaultunseededtierqueueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedDuoQueueDefaultUnseededTierQueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedduoqueuedefaultunseededtierqueueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedDuoQueueDefaultUnseededTierQueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedduoqueuedefaultunseededtierqueueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedDuoQueueDefaultUnseededTierQueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedduoqueuerestrictionmaxdelta() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedDuoQueueRestrictionMaxDelta"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedduoqueuerestrictionmode() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedDuoQueueRestrictionMode"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighmmrpremademaxsize() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighMmrPremadeMaxSize"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighmmrpremaderank() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighMmrPremadeRank"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighmmrpremaderestrictionenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighMmrPremadeRestrictionEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighmmrpremadetier() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighMmrPremadeTier"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskilllowesttier() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillLowestTier"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskilllowesttierqueueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillLowestTierQueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskilllowesttierqueueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillLowestTierQueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskilllowesttierqueueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillLowestTierQueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskilllowesttierqueueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillLowestTierQueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskilllowesttierqueueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillLowestTierQueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmaxdelta() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionMaxDelta"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmaxdeltaqueueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionMaxDeltaQueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmaxdeltaqueueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionMaxDeltaQueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmaxdeltaqueueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionMaxDeltaQueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmaxdeltaqueueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionMaxDeltaQueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmaxdeltaqueueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionMaxDeltaQueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmode() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionMode"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmodequeueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionModeQueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmodequeueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionModeQueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmodequeueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionModeQueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmodequeueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionModeQueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedhighskillrestrictionmodequeueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedHighSkillRestrictionModeQueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmaxdelta() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionMaxDelta"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmaxdeltaqueueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionMaxDeltaQueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmaxdeltaqueueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionMaxDeltaQueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmaxdeltaqueueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionMaxDeltaQueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmaxdeltaqueueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionMaxDeltaQueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmaxdeltaqueueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionMaxDeltaQueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmode() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionMode"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmodequeueid410() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionModeQueueId410"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmodequeueid420() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionModeQueueId420"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmodequeueid440() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionModeQueueId440"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmodequeueid470() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionModeQueueId470"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionRankedlowskillrestrictionmodequeueid9() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_RankedLowSkillRestrictionModeQueueId9"
}
func (EventTopic) PlatformConfigV1NamespacesQueuerestrictionServiceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_QueueRestriction_ServiceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesRegalia() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Regalia"
}
func (EventTopic) PlatformConfigV1NamespacesReplays() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Replays"
}
func (EventTopic) PlatformConfigV1NamespacesReplaysMinsupportedgameserverversion() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Replays_MinSupportedGameServerVersion"
}
func (EventTopic) PlatformConfigV1NamespacesReplaysThirdpersonaccessiblegamequeues() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Replays_ThirdPersonAccessibleGameQueues"
}
func (EventTopic) PlatformConfigV1NamespacesRewards() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Rewards"
}
func (EventTopic) PlatformConfigV1NamespacesRewardsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Rewards_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesSlcuchampionselecttring() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SLcuChampionSelecttring"
}
func (EventTopic) PlatformConfigV1NamespacesSlcuchampionselecttringRandomchampionratelimitinterval() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SLcuChampionSelecttring_RandomChampionRateLimitInterval"
}
func (EventTopic) PlatformConfigV1NamespacesSlcuchampionselecttringRandomchampionratelimitmaxactions() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SLcuChampionSelecttring_RandomChampionRateLimitMaxActions"
}
func (EventTopic) PlatformConfigV1NamespacesSanitizer() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Sanitizer"
}
func (EventTopic) PlatformConfigV1NamespacesSanitizerLevel1unfilter() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Sanitizer_Level1Unfilter"
}
func (EventTopic) PlatformConfigV1NamespacesSeasonreward() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SeasonReward"
}
func (EventTopic) PlatformConfigV1NamespacesSeasonrewardEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SeasonReward_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesSeasonrewardMaximumTeamRewardLevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SeasonReward_Maximum_team_reward_level"
}
func (EventTopic) PlatformConfigV1NamespacesSeasonrewardMinimumPointsPerRewardLevel() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SeasonReward_Minimum_points_per_reward_level"
}
func (EventTopic) PlatformConfigV1NamespacesSeasonrewardMinimumWinTeamRewardLevel1() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SeasonReward_Minimum_win_team_reward_level_1"
}
func (EventTopic) PlatformConfigV1NamespacesSeasonrewardMinimumWinTeamRewardLevel2() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SeasonReward_Minimum_win_team_reward_level_2"
}
func (EventTopic) PlatformConfigV1NamespacesSeasonrewardMinimumWinTeamRewardLevel3() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SeasonReward_Minimum_win_team_reward_level_3"
}
func (EventTopic) PlatformConfigV1NamespacesSeasonrewardQualificationwarningenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SeasonReward_QualificationWarningEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesSeasonrewardServicecallenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SeasonReward_ServiceCallEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesSentry() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Sentry"
}
func (EventTopic) PlatformConfigV1NamespacesServicestatusapi() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ServiceStatusAPI"
}
func (EventTopic) PlatformConfigV1NamespacesServicestatusapiEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ServiceStatusAPI_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesSharematchhistory() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ShareMatchHistory"
}
func (EventTopic) PlatformConfigV1NamespacesSharematchhistoryAdvancedgamedetailsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ShareMatchHistory_AdvancedGameDetailsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesSharematchhistoryAdvancedgamedetailsurltemplate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ShareMatchHistory_AdvancedGameDetailsUrlTemplate"
}
func (EventTopic) PlatformConfigV1NamespacesSharematchhistoryMatchdetailsurltemplate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ShareMatchHistory_MatchDetailsUrlTemplate"
}
func (EventTopic) PlatformConfigV1NamespacesSharematchhistoryMatchhistoryenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ShareMatchHistory_MatchHistoryEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesSharematchhistoryMatchhistoryurltemplate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ShareMatchHistory_MatchHistoryUrlTemplate"
}
func (EventTopic) PlatformConfigV1NamespacesSharematchhistoryShareendofgameenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ShareMatchHistory_ShareEndOfGameEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesSharematchhistorySharegameurltemplate() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ShareMatchHistory_ShareGameUrlTemplate"
}
func (EventTopic) PlatformConfigV1NamespacesSkinrentals() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SkinRentals"
}
func (EventTopic) PlatformConfigV1NamespacesSkinrentalsEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SkinRentals_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesSkinsviewer() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SkinsViewer"
}
func (EventTopic) PlatformConfigV1NamespacesSkinsviewerVintageskinsummonericonconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SkinsViewer_VintageSkinSummonerIconConfig"
}
func (EventTopic) PlatformConfigV1NamespacesSocialleaderboard() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SocialLeaderboard"
}
func (EventTopic) PlatformConfigV1NamespacesSocialleaderboardIssocialleaderboardenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SocialLeaderboard_IsSocialLeaderboardEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesSocialleaderboardMinstillcacheexpiry() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SocialLeaderboard_MinsTillCacheExpiry"
}
func (EventTopic) PlatformConfigV1NamespacesSocialleaderboardSecstillavailabilitycacheexpiry() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SocialLeaderboard_SecsTillAvailabilityCacheExpiry"
}
func (EventTopic) PlatformConfigV1NamespacesSocialleaderboardUsesocialleaderboardleaguesendpoint() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SocialLeaderboard_UseSocialLeaderboardLeaguesEndpoint"
}
func (EventTopic) PlatformConfigV1NamespacesSpectator() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Spectator"
}
func (EventTopic) PlatformConfigV1NamespacesSpectatorSavereconnectinfoenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Spectator_SaveReconnectInfoEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesString() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_String"
}
func (EventTopic) PlatformConfigV1NamespacesStringString() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_String_String"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayers() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayersEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayersFriendsoffriendsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers_FriendsOfFriendsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayersFriendsoffriendslimit() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers_FriendsOfFriendsLimit"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayersHonoredplayerslimit() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers_HonoredPlayersLimit"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayersMaxnumreplacements() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers_MaxNumReplacements"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayersMaxnumsuggestedplayers() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers_MaxNumSuggestedPlayers"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayersOnlinefriendslimit() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers_OnlineFriendsLimit"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayersPreviouspremadeslimit() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers_PreviousPremadesLimit"
}
func (EventTopic) PlatformConfigV1NamespacesSuggestedplayersVictoriouscomradeslimit() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_SuggestedPlayers_VictoriousComradesLimit"
}
func (EventTopic) PlatformConfigV1NamespacesSummoner() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Summoner"
}
func (EventTopic) PlatformConfigV1NamespacesSummonerConfigrefreshintervalseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Summoner_ConfigRefreshIntervalSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesSummonerJwtenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Summoner_JWTEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesSummonerJwtmaxtimeoutseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Summoner_JWTMaxTimeoutSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesSummonerJwtmintimeoutseconds() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Summoner_JWTMinTimeoutSeconds"
}
func (EventTopic) PlatformConfigV1NamespacesSummonerSummonerprofilecacheenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Summoner_SummonerProfileCacheEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTeamboost() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBoost"
}
func (EventTopic) PlatformConfigV1NamespacesTeamboostAllskinenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBoost_AllSkinEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTeamboostRandomskinenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBoost_RandomSkinEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraft() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftEnablechampionselectpreferences() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_EnableChampionSelectPreferences"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftEstimatedwaitadjustmentenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_EstimatedWaitAdjustmentEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftFillprimarywaitfactor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_FillPrimaryWaitFactor"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftFillsecondarywaitfactor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_FillSecondaryWaitFactor"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftLogalllcdsmessages() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_LogAllLCDSMessages"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftMidprimarywaitfactor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_MidPrimaryWaitFactor"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftSendafkcheckmetricsenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_SendAfkCheckMetricsEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftSendsummonericonjwt() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_SendSummonerIconJwt"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftServicecalltimeoutmillis() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_ServiceCallTimeoutMillis"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftSkinpurchaseenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_SkinPurchaseEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftSupportprimarywaitfactor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_SupportPrimaryWaitFactor"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftSupportsecondarywaitfactor() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_SupportSecondaryWaitFactor"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftTbrerollserviceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_TBRerollServiceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftTakeoverenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_TakeoverEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTeambuilderdraftUnlocklockinbuttontimeoutenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TeamBuilderDraft_UnlockLockInButtonTimeoutEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTencentantiaddictionEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_TencentAntiAddiction_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesThirdpartyverification() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ThirdPartyVerification"
}
func (EventTopic) PlatformConfigV1NamespacesThirdpartyverificationCahserviceenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_ThirdPartyVerification_CAHServiceEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesTrophies() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Trophies"
}
func (EventTopic) PlatformConfigV1NamespacesTrophiesIsenabledonprofile() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Trophies_IsEnabledOnProfile"
}
func (EventTopic) PlatformConfigV1NamespacesTrophiesIsothersummonersprofileenabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Trophies_IsOtherSummonersProfileEnabled"
}
func (EventTopic) PlatformConfigV1NamespacesVoice() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Voice"
}
func (EventTopic) PlatformConfigV1NamespacesVoiceEnabled() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_Voice_Enabled"
}
func (EventTopic) PlatformConfigV1NamespacesWardskinconfig() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_WardSkinConfig"
}
func (EventTopic) PlatformConfigV1NamespacesWardskinconfigUseloadouts() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_WardSkinConfig_UseLoadouts"
}
func (EventTopic) PlatformConfigV1NamespacesWardskinconfigWardskinselection() string {
	return "OnJsonApiEvent_lol-platform-config_v1_namespaces_WardSkinConfig_WardSkinSelection"
}

// Lol Player Behavior events
func (EventTopic) PlayerBehaviorV1Config() string {
	return "OnJsonApiEvent_lol-player-behavior_v1_config"
}
func (EventTopic) PlayerBehaviorV2ReporterFeedback() string {
	return "OnJsonApiEvent_lol-player-behavior_v2_reporter-feedback"
}
func (EventTopic) PlayerBehaviorV3ReformCards() string {
	return "OnJsonApiEvent_lol-player-behavior_v3_reform-cards"
}

// Lol Player Preferences events
func (EventTopic) PlayerPreferencesV1Ready() string {
	return "OnJsonApiEvent_lol-player-preferences_v1_player-preferences-ready"
}

// Lol Pre End Of Game events
func (EventTopic) PreEndOfGameV1Currentsequenceevent() string {
	return "OnJsonApiEvent_lol-pre-end-of-game_v1_currentSequenceEvent"
}

// Lol Premade Voice events
func (EventTopic) PremadeVoiceV1Availability() string {
	return "OnJsonApiEvent_lol-premade-voice_v1_availability"
}
func (EventTopic) PremadeVoiceV1Capturedevices() string {
	return "OnJsonApiEvent_lol-premade-voice_v1_capturedevices"
}
func (EventTopic) PremadeVoiceV1FirstExperience() string {
	return "OnJsonApiEvent_lol-premade-voice_v1_first-experience"
}
func (EventTopic) PremadeVoiceV1ParticipantRecords() string {
	return "OnJsonApiEvent_lol-premade-voice_v1_participant-records"
}
func (EventTopic) PremadeVoiceV1Settings() string {
	return "OnJsonApiEvent_lol-premade-voice_v1_settings"
}

// Lol Progression events
func (EventTopic) ProgressionV1Groups() string { return "OnJsonApiEvent_lol-progression_v1_groups" }
func (EventTopic) ProgressionV1Ready() string  { return "OnJsonApiEvent_lol-progression_v1_ready" }

// Lol Publishing Content events
func (EventTopic) PublishingContentV1Listeners() string {
	return "OnJsonApiEvent_lol-publishing-content_v1_listeners"
}
func (EventTopic) PublishingContentV1Ready() string {
	return "OnJsonApiEvent_lol-publishing-content_v1_ready"
}
func (EventTopic) PublishingContentV1Settings() string {
	return "OnJsonApiEvent_lol-publishing-content_v1_settings"
}

// Lol Purchase Widget events
func (EventTopic) PurchaseWidgetV1Configuration() string {
	return "OnJsonApiEvent_lol-purchase-widget_v1_configuration"
}
func (EventTopic) PurchaseWidgetV3PurchaseOfferOrderStatuses() string {
	return "OnJsonApiEvent_lol-purchase-widget_v3_purchase-offer-order-statuses"
}

// Lol Ranked events
func (EventTopic) RankedV1ChallengerLaddersEnabled() string {
	return "OnJsonApiEvent_lol-ranked_v1_challenger-ladders-enabled"
}
func (EventTopic) RankedV1CurrentRankedStats() string {
	return "OnJsonApiEvent_lol-ranked_v1_current-ranked-stats"
}
func (EventTopic) RankedV1GlobalNotifications() string {
	return "OnJsonApiEvent_lol-ranked_v1_global-notifications"
}
func (EventTopic) RankedV1Notifications() string { return "OnJsonApiEvent_lol-ranked_v1_notifications" }
func (EventTopic) RankedV1Stats() string         { return "OnJsonApiEvent_lol-ranked_v1_ranked-stats" }
func (EventTopic) RankedV1SignedRankedStats() string {
	return "OnJsonApiEvent_lol-ranked_v1_signed-ranked-stats"
}
func (EventTopic) RankedV1TopRatedLaddersEnabled() string {
	return "OnJsonApiEvent_lol-ranked_v1_top-rated-ladders-enabled"
}

// Lol Regalia events
func (EventTopic) RegaliaV2Config() string    { return "OnJsonApiEvent_lol-regalia_v2_config" }
func (EventTopic) RegaliaV2Summoners() string { return "OnJsonApiEvent_lol-regalia_v2_summoners" }

// Lol Remedy events
func (EventTopic) RemedyV1Notifications() string {
	return "OnJsonApiEvent_lol-remedy_v1_remedy-notifications"
}

// Lol Replays events
func (EventTopic) ReplaysV1Configuration() string {
	return "OnJsonApiEvent_lol-replays_v1_configuration"
}
func (EventTopic) ReplaysV1Metadata() string { return "OnJsonApiEvent_lol-replays_v1_metadata" }
func (EventTopic) ReplaysV1Rofls() string    { return "OnJsonApiEvent_lol-replays_v1_rofls" }

// Lol Rewards events
func (EventTopic) RewardsV1Grants() string { return "OnJsonApiEvent_lol-rewards_v1_grants" }
func (EventTopic) RewardsV1Groups() string { return "OnJsonApiEvent_lol-rewards_v1_groups" }

// Lol Rso Auth events
func (EventTopic) LoLRsoAuthConfigurationV3() string {
	return "OnJsonApiEvent_lol-rso-auth_configuration_v3"
}
func (EventTopic) LoLRsoAuthV1Authorization() string {
	return "OnJsonApiEvent_lol-rso-auth_v1_authorization"
}

// Lol Settings events
func (EventTopic) SettingsV1Account() string { return "OnJsonApiEvent_lol-settings_v1_account" }
func (EventTopic) SettingsV1Local() string   { return "OnJsonApiEvent_lol-settings_v1_local" }
func (EventTopic) SettingsV2Account() string { return "OnJsonApiEvent_lol-settings_v2_account" }
func (EventTopic) SettingsV2Config() string  { return "OnJsonApiEvent_lol-settings_v2_config" }
func (EventTopic) SettingsV2Local() string   { return "OnJsonApiEvent_lol-settings_v2_local" }
func (EventTopic) SettingsV2Ready() string   { return "OnJsonApiEvent_lol-settings_v2_ready" }

// Lol Shoppefront events
func (EventTopic) ShoppefrontV1Ready() string  { return "OnJsonApiEvent_lol-shoppefront_v1_ready" }
func (EventTopic) ShoppefrontV1Stores() string { return "OnJsonApiEvent_lol-shoppefront_v1_stores" }

// Lol Shutdown events
func (EventTopic) ShutdownV1Notification() string {
	return "OnJsonApiEvent_lol-shutdown_v1_notification"
}

// Lol Spectator events
func (EventTopic) SpectatorV1Spectate() string { return "OnJsonApiEvent_lol-spectator_v1_spectate" }

// Lol Statstones events
func (EventTopic) StatstonesV2PlayerSummarySelf() string {
	return "OnJsonApiEvent_lol-statstones_v2_player-summary-self"
}

// Lol Store events
func (EventTopic) StoreV1Getstoreurl() string { return "OnJsonApiEvent_lol-store_v1_getStoreUrl" }
func (EventTopic) StoreV1Ready() string       { return "OnJsonApiEvent_lol-store_v1_store-ready" }

// Lol Suggested Players events
func (EventTopic) SuggestedPlayersV1() string {
	return "OnJsonApiEvent_lol-suggested-players_v1_suggested-players"
}

// Lol Summoner events
func (EventTopic) SummonerV1CurrentSummoner() string {
	return "OnJsonApiEvent_lol-summoner_v1_current-summoner"
}
func (EventTopic) SummonerV1PlayerAliasState() string {
	return "OnJsonApiEvent_lol-summoner_v1_player-alias-state"
}
func (EventTopic) SummonerV1PlayerNameMode() string {
	return "OnJsonApiEvent_lol-summoner_v1_player-name-mode"
}
func (EventTopic) SummonerV1Status() string { return "OnJsonApiEvent_lol-summoner_v1_status" }
func (EventTopic) SummonerV1RequestsReady() string {
	return "OnJsonApiEvent_lol-summoner_v1_summoner-requests-ready"
}

// Lol Summoner Profiles events
func (EventTopic) SummonerProfilesV1GetSummonerLevelView() string {
	return "OnJsonApiEvent_lol-summoner-profiles_v1_get-summoner-level-view"
}

// Lol Tastes events
func (EventTopic) TastesV1Ready() string { return "OnJsonApiEvent_lol-tastes_v1_ready" }

// Lol Tft events
func (EventTopic) TftV1() string { return "OnJsonApiEvent_lol-tft_v1_tft" }

// Lol Tft Pass events
func (EventTopic) TftPassV1ActivePasses() string {
	return "OnJsonApiEvent_lol-tft-pass_v1_active-passes"
}
func (EventTopic) TftPassV1BattlePass() string { return "OnJsonApiEvent_lol-tft-pass_v1_battle-pass" }
func (EventTopic) TftPassV1ConfigFetched() string {
	return "OnJsonApiEvent_lol-tft-pass_v1_config-fetched"
}
func (EventTopic) TftPassV1EventPass() string { return "OnJsonApiEvent_lol-tft-pass_v1_event-pass" }
func (EventTopic) TftPassV1Ready() string     { return "OnJsonApiEvent_lol-tft-pass_v1_ready" }
func (EventTopic) TftPassV1SkillTreePass() string {
	return "OnJsonApiEvent_lol-tft-pass_v1_skill-tree-pass"
}

// Lol Tft Skill Tree events
func (EventTopic) TftSkillTreeV1SkillTree() string {
	return "OnJsonApiEvent_lol-tft-skill-tree_v1_skill-tree"
}

// Lol Tft Team Planner events
func (EventTopic) TftTeamPlannerV1Config() string {
	return "OnJsonApiEvent_lol-tft-team-planner_v1_config"
}
func (EventTopic) TftTeamPlannerV1Ftue() string { return "OnJsonApiEvent_lol-tft-team-planner_v1_ftue" }
func (EventTopic) TftTeamPlannerV1PreviousContext() string {
	return "OnJsonApiEvent_lol-tft-team-planner_v1_previous-context"
}
func (EventTopic) TftTeamPlannerV1Set() string { return "OnJsonApiEvent_lol-tft-team-planner_v1_set" }
func (EventTopic) TftTeamPlannerV2Reminders() string {
	return "OnJsonApiEvent_lol-tft-team-planner_v2_reminders"
}

// Lol Tft Troves events
func (EventTopic) TftTrovesV1Config() string { return "OnJsonApiEvent_lol-tft-troves_v1_config" }
func (EventTopic) TftTrovesV1MilestonesGroupId() string {
	return "OnJsonApiEvent_lol-tft-troves_v1_milestones-group-id"
}

// Lol Yourshop events
func (EventTopic) YourshopV1Modal() string  { return "OnJsonApiEvent_lol-yourshop_v1_modal" }
func (EventTopic) YourshopV1Ready() string  { return "OnJsonApiEvent_lol-yourshop_v1_ready" }
func (EventTopic) YourshopV1Status() string { return "OnJsonApiEvent_lol-yourshop_v1_status" }

// Loyalty events
func (EventTopic) LoyaltyV1Resource() string { return "OnJsonApiEvent_loyalty_v1_loyalty-resource" }

// Mailbox events
func (EventTopic) MailboxV1CheckNewMail() string { return "OnJsonApiEvent_mailbox_v1_check-new-mail" }

// Memory events
func (EventTopic) MemoryV1FeProcessesReady() string {
	return "OnJsonApiEvent_memory_v1_fe-processes-ready"
}

// Patcher events
func (EventTopic) PatcherV1Notifications() string { return "OnJsonApiEvent_patcher_v1_notifications" }
func (EventTopic) PatcherV1Products() string      { return "OnJsonApiEvent_patcher_v1_products" }
func (EventTopic) PatcherV1Status() string        { return "OnJsonApiEvent_patcher_v1_status" }

// Player Account events
func (EventTopic) PlayerAccountV1Aliases() string { return "OnJsonApiEvent_player-account_aliases_v1" }

// Player Notifications events
func (EventTopic) PlayerNotificationsV1Notifications() string {
	return "OnJsonApiEvent_player-notifications_v1_notifications"
}

// Player Reporting events
func (EventTopic) PlayerReportingV1ReporterFeedback() string {
	return "OnJsonApiEvent_player-reporting_v1_reporter-feedback"
}

// Plugin Manager events
func (EventTopic) PluginManagerV1ExternalPlugins() string {
	return "OnJsonApiEvent_plugin-manager_v1_external-plugins"
}
func (EventTopic) PluginManagerV1Status() string { return "OnJsonApiEvent_plugin-manager_v1_status" }

// Process Control events
func (EventTopic) ProcessControlV1Process() string {
	return "OnJsonApiEvent_process-control_v1_process"
}

// Product Integration events
func (EventTopic) ProductIntegrationV1AppUpdate() string {
	return "OnJsonApiEvent_product-integration_v1_app-update"
}

// Product Metadata events
func (EventTopic) ProductMetadataV2Products() string {
	return "OnJsonApiEvent_product-metadata_v2_products"
}

// Product Session events
func (EventTopic) ProductSessionV1ExternalSessions() string {
	return "OnJsonApiEvent_product-session_v1_external-sessions"
}

// Riot Messaging Service events
func (EventTopic) RiotMessagingServiceV1Message() string {
	return "OnJsonApiEvent_riot-messaging-service_v1_message"
}
func (EventTopic) RiotMessagingServiceV1OutOfSync() string {
	return "OnJsonApiEvent_riot-messaging-service_v1_out-of-sync"
}
func (EventTopic) RiotMessagingServiceV1Session() string {
	return "OnJsonApiEvent_riot-messaging-service_v1_session"
}
func (EventTopic) RiotMessagingServiceV1State() string {
	return "OnJsonApiEvent_riot-messaging-service_v1_state"
}

// Riotclient events
func (EventTopic) RiotclientAffinity() string { return "OnJsonApiEvent_riotclient_affinity" }
func (EventTopic) RiotclientAppPort() string  { return "OnJsonApiEvent_riotclient_app-port" }
func (EventTopic) RiotclientNewArgs() string  { return "OnJsonApiEvent_riotclient_new-args" }
func (EventTopic) RiotclientPreShutdownBegin() string {
	return "OnJsonApiEvent_riotclient_pre-shutdown_begin"
}
func (EventTopic) RiotclientRegionLocale() string { return "OnJsonApiEvent_riotclient_region-locale" }
func (EventTopic) RiotclientV1SystemInfo() string { return "OnJsonApiEvent_riotclient_system-info_v1" }
func (EventTopic) RiotclientUxCrashCount() string { return "OnJsonApiEvent_riotclient_ux-crash-count" }
func (EventTopic) RiotclientUxStateRequest() string {
	return "OnJsonApiEvent_riotclient_ux-state_request"
}
func (EventTopic) RiotclientV1CrashReporting() string {
	return "OnJsonApiEvent_riotclient_v1_crash-reporting"
}
func (EventTopic) RiotclientZoomScale() string { return "OnJsonApiEvent_riotclient_zoom-scale" }

// Rso Auth events
func (EventTopic) RsoAuthConfigurationV3() string { return "OnJsonApiEvent_rso-auth_configuration_v3" }
func (EventTopic) RsoAuthV1Session() string       { return "OnJsonApiEvent_rso-auth_v1_session" }
func (EventTopic) RsoAuthV1Userinfo() string      { return "OnJsonApiEvent_rso-auth_v1_userinfo" }
func (EventTopic) RsoAuthV2Authorizations() string {
	return "OnJsonApiEvent_rso-auth_v2_authorizations"
}

// Sanitizer events
func (EventTopic) SanitizerV1Status() string { return "OnJsonApiEvent_sanitizer_v1_status" }

// Scd events
func (EventTopic) ScdV1Cookies() string { return "OnJsonApiEvent_scd_v1_cookies" }

// System events
func (EventTopic) SystemV1Builds() string { return "OnJsonApiEvent_system_v1_builds" }

// Voice Chat events
func (EventTopic) VoiceChatV1AudioProperties() string {
	return "OnJsonApiEvent_voice-chat_v1_audio-properties"
}
func (EventTopic) VoiceChatV2Devices() string  { return "OnJsonApiEvent_voice-chat_v2_devices" }
func (EventTopic) VoiceChatV2State() string    { return "OnJsonApiEvent_voice-chat_v2_state" }
func (EventTopic) VoiceChatV3Sessions() string { return "OnJsonApiEvent_voice-chat_v3_sessions" }
func (EventTopic) VoiceChatV3Settings() string { return "OnJsonApiEvent_voice-chat_v3_settings" }

// LCDS events
func (EventTopic) LcdsBroadcastnotification() string {
	return "OnLcdsEvent_com_riotgames_platform_broadcast_BroadcastNotification"
}
func (EventTopic) LcdsClientdynamicconfigurationnotification() string {
	return "OnLcdsEvent_com_riotgames_platform_client_dynamic_configuration_ClientDynamicConfigurationNotification"
}
func (EventTopic) LcdsGamedto() string { return "OnLcdsEvent_com_riotgames_platform_game_GameDTO" }
func (EventTopic) LcdsPlayercredentialsdto() string {
	return "OnLcdsEvent_com_riotgames_platform_game_PlayerCredentialsDto"
}
func (EventTopic) LcdsTeamskinrentaldto() string {
	return "OnLcdsEvent_com_riotgames_platform_game_TeamSkinRentalDTO"
}
func (EventTopic) LcdsGamenotification() string {
	return "OnLcdsEvent_com_riotgames_platform_game_message_GameNotification"
}
func (EventTopic) LcdsInvitationrequest() string {
	return "OnLcdsEvent_com_riotgames_platform_gameinvite_contract_InvitationRequest"
}
func (EventTopic) LcdsInviteprivileges() string {
	return "OnLcdsEvent_com_riotgames_platform_gameinvite_contract_InvitePrivileges"
}
func (EventTopic) LcdsLobbystatus() string {
	return "OnLcdsEvent_com_riotgames_platform_gameinvite_contract_LobbyStatus"
}
func (EventTopic) LcdsRemovedfromlobbynotification() string {
	return "OnLcdsEvent_com_riotgames_platform_gameinvite_contract_RemovedFromLobbyNotification"
}
func (EventTopic) LcdsLeaverbusterlowpriorityqueueabandoned() string {
	return "OnLcdsEvent_com_riotgames_platform_leaverbuster_event_messaging_LeaverBusterLowPriorityQueueAbandoned"
}
func (EventTopic) LcdsLoyaltystatechangenotification() string {
	return "OnLcdsEvent_com_riotgames_platform_loyalty_message_LoyaltyStateChangeNotification"
}
func (EventTopic) LcdsSearchingformatchnotification() string {
	return "OnLcdsEvent_com_riotgames_platform_matchmaking_SearchingForMatchNotification"
}
func (EventTopic) LcdsForcedclientshutdown() string {
	return "OnLcdsEvent_com_riotgames_platform_messaging_ForcedClientShutdown"
}
func (EventTopic) LcdsRentalupdatenotification() string {
	return "OnLcdsEvent_com_riotgames_platform_messaging_RentalUpdateNotification"
}
func (EventTopic) LcdsStorefulfillmentnotification() string {
	return "OnLcdsEvent_com_riotgames_platform_messaging_StoreFulfillmentNotification"
}
func (EventTopic) LcdsSimpledialogmessage() string {
	return "OnLcdsEvent_com_riotgames_platform_messaging_persistence_SimpleDialogMessage"
}
func (EventTopic) LcdsClientsystemstatesnotification() string {
	return "OnLcdsEvent_com_riotgames_platform_systemstate_ClientSystemStatesNotification"
}
func (EventTopic) LcdsTradecontractdto() string {
	return "OnLcdsEvent_com_riotgames_platform_trade_api_contract_TradeContractDTO"
}

// Other events
func (EventTopic) RegionLocaleChanged() string { return "OnRegionLocaleChanged" }

// Fired when an async service proxy message is received.
func (EventTopic) ServiceProxyAsyncEvent() string  { return "OnServiceProxyAsyncEvent" }
func (EventTopic) ServiceProxyMethodEvent() string { return "OnServiceProxyMethodEvent" }

// Fired when a uuid-based service proxy message is received.
func (EventTopic) ServiceProxyUuidEvent() string { return "OnServiceProxyUuidEvent" }
