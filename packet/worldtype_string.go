// Code generated by "stringer -type=WorldType"; DO NOT EDIT

package packet

import "fmt"

const _WorldType_name = "CMSG_BOOTMECMSG_DBLOOKUPSMSG_DBLOOKUPCMSG_QUERY_OBJECT_POSITIONSMSG_QUERY_OBJECT_POSITIONCMSG_QUERY_OBJECT_ROTATIONSMSG_QUERY_OBJECT_ROTATIONCMSG_WORLD_TELEPORTCMSG_TELEPORT_TO_UNITCMSG_ZONE_MAPSMSG_ZONE_MAPCMSG_DEBUG_CHANGECELLZONECMSG_MOVE_CHARACTER_CHEATSMSG_MOVE_CHARACTER_CHEATCMSG_RECHARGECMSG_LEARN_SPELLCMSG_CREATEMONSTERCMSG_DESTROYMONSTERCMSG_CREATEITEMCMSG_CREATEGAMEOBJECTSMSG_CHECK_FOR_BOTSCMSG_MAKEMONSTERATTACKGUIDCMSG_BOT_DETECTED2CMSG_FORCEACTIONCMSG_FORCEACTIONONOTHERCMSG_FORCEACTIONSHOWSMSG_FORCEACTIONSHOWCMSG_PETGODMODESMSG_PETGODMODESMSG_REFER_A_FRIEND_EXPIREDCMSG_WEATHER_SPEED_CHEATCMSG_UNDRESSPLAYERCMSG_BEASTMASTERCMSG_GODMODESMSG_GODMODECMSG_CHEAT_SETMONEYCMSG_LEVEL_CHEATCMSG_PET_LEVEL_CHEATCMSG_SET_WORLDSTATECMSG_COOLDOWN_CHEATCMSG_USE_SKILL_CHEATCMSG_FLAG_QUESTCMSG_FLAG_QUEST_FINISHCMSG_CLEAR_QUESTCMSG_SEND_EVENTCMSG_DEBUG_AISTATESMSG_DEBUG_AISTATECMSG_DISABLE_PVP_CHEATCMSG_ADVANCE_SPAWN_TIMESMSG_DESTRUCTIBLE_BUILDING_DAMAGECMSG_AUTH_SRP6_BEGINCMSG_AUTH_SRP6_PROOFCMSG_AUTH_SRP6_RECODECMSG_CHAR_CREATECMSG_CHAR_ENUMCMSG_CHAR_DELETESMSG_AUTH_SRP6_RESPONSESMSG_CHAR_CREATESMSG_CHAR_ENUMSMSG_CHAR_DELETECMSG_PLAYER_LOGINSMSG_NEW_WORLDSMSG_TRANSFER_PENDINGSMSG_TRANSFER_ABORTEDSMSG_CHARACTER_LOGIN_FAILEDSMSG_LOGIN_SETTIMESPEEDSMSG_GAMETIME_UPDATECMSG_GAMETIME_SETSMSG_GAMETIME_SETCMSG_GAMESPEED_SETSMSG_GAMESPEED_SETCMSG_SERVERTIMESMSG_SERVERTIMECMSG_PLAYER_LOGOUTCMSG_LOGOUT_REQUESTSMSG_LOGOUT_RESPONSESMSG_LOGOUT_COMPLETECMSG_LOGOUT_CANCELSMSG_LOGOUT_CANCEL_ACKCMSG_NAME_QUERYSMSG_NAME_QUERY_RESPONSECMSG_PET_NAME_QUERYSMSG_PET_NAME_QUERY_RESPONSECMSG_GUILD_QUERYSMSG_GUILD_QUERY_RESPONSECMSG_ITEM_QUERY_SINGLECMSG_ITEM_QUERY_MULTIPLESMSG_ITEM_QUERY_SINGLE_RESPONSESMSG_ITEM_QUERY_MULTIPLE_RESPONSECMSG_PAGE_TEXT_QUERYSMSG_PAGE_TEXT_QUERY_RESPONSECMSG_QUEST_QUERYSMSG_QUEST_QUERY_RESPONSECMSG_GAMEOBJECT_QUERYSMSG_GAMEOBJECT_QUERY_RESPONSECMSG_CREATURE_QUERYSMSG_CREATURE_QUERY_RESPONSECMSG_WHOSMSG_WHOCMSG_WHOISSMSG_WHOISCMSG_CONTACT_LISTSMSG_CONTACT_LISTSMSG_FRIEND_STATUSCMSG_ADD_FRIENDCMSG_DEL_FRIENDCMSG_SET_CONTACT_NOTESCMSG_ADD_IGNORECMSG_DEL_IGNORECMSG_GROUP_INVITESMSG_GROUP_INVITECMSG_GROUP_CANCELSMSG_GROUP_CANCELCMSG_GROUP_ACCEPTCMSG_GROUP_DECLINESMSG_GROUP_DECLINECMSG_GROUP_UNINVITECMSG_GROUP_UNINVITE_GUIDSMSG_GROUP_UNINVITECMSG_GROUP_SET_LEADERSMSG_GROUP_SET_LEADERCMSG_LOOT_METHODCMSG_GROUP_DISBANDSMSG_GROUP_DESTROYEDSMSG_GROUP_LISTSMSG_PARTY_MEMBER_STATSSMSG_PARTY_COMMAND_RESULTUMSG_UPDATE_GROUP_MEMBERSCMSG_GUILD_CREATECMSG_GUILD_INVITESMSG_GUILD_INVITECMSG_GUILD_ACCEPTCMSG_GUILD_DECLINESMSG_GUILD_DECLINECMSG_GUILD_INFOSMSG_GUILD_INFOCMSG_GUILD_ROSTERSMSG_GUILD_ROSTERCMSG_GUILD_PROMOTECMSG_GUILD_DEMOTECMSG_GUILD_LEAVECMSG_GUILD_REMOVECMSG_GUILD_DISBANDCMSG_GUILD_LEADERCMSG_GUILD_MOTDSMSG_GUILD_EVENTSMSG_GUILD_COMMAND_RESULTUMSG_UPDATE_GUILDCMSG_MESSAGECHATSMSG_MESSAGECHATCMSG_JOIN_CHANNELCMSG_LEAVE_CHANNELSMSG_CHANNEL_NOTIFYCMSG_CHANNEL_LISTSMSG_CHANNEL_LISTCMSG_CHANNEL_PASSWORDCMSG_CHANNEL_SET_OWNERCMSG_CHANNEL_OWNERCMSG_CHANNEL_MODERATORCMSG_CHANNEL_UNMODERATORCMSG_CHANNEL_MUTECMSG_CHANNEL_UNMUTECMSG_CHANNEL_INVITECMSG_CHANNEL_KICKCMSG_CHANNEL_BANCMSG_CHANNEL_UNBANCMSG_CHANNEL_ANNOUNCEMENTSCMSG_CHANNEL_MODERATESMSG_UPDATE_OBJECTSMSG_DESTROY_OBJECTCMSG_USE_ITEMCMSG_OPEN_ITEMCMSG_READ_ITEMSMSG_READ_ITEM_OKSMSG_READ_ITEM_FAILEDSMSG_ITEM_COOLDOWNCMSG_GAMEOBJ_USECMSG_DESTROY_ITEMSSMSG_GAMEOBJECT_CUSTOM_ANIMCMSG_AREATRIGGERMSG_MOVE_START_FORWARDMSG_MOVE_START_BACKWARDMSG_MOVE_STOPMSG_MOVE_START_STRAFE_LEFTMSG_MOVE_START_STRAFE_RIGHTMSG_MOVE_STOP_STRAFEMSG_MOVE_JUMPMSG_MOVE_START_TURN_LEFTMSG_MOVE_START_TURN_RIGHTMSG_MOVE_STOP_TURNMSG_MOVE_START_PITCH_UPMSG_MOVE_START_PITCH_DOWNMSG_MOVE_STOP_PITCHMSG_MOVE_SET_RUN_MODEMSG_MOVE_SET_WALK_MODEMSG_MOVE_TOGGLE_LOGGINGMSG_MOVE_TELEPORTMSG_MOVE_TELEPORT_CHEATMSG_MOVE_TELEPORT_ACKMSG_MOVE_TOGGLE_FALL_LOGGINGMSG_MOVE_FALL_LANDMSG_MOVE_START_SWIMMSG_MOVE_STOP_SWIMMSG_MOVE_SET_RUN_SPEED_CHEATMSG_MOVE_SET_RUN_SPEEDMSG_MOVE_SET_RUN_BACK_SPEED_CHEATMSG_MOVE_SET_RUN_BACK_SPEEDMSG_MOVE_SET_WALK_SPEED_CHEATMSG_MOVE_SET_WALK_SPEEDMSG_MOVE_SET_SWIM_SPEED_CHEATMSG_MOVE_SET_SWIM_SPEEDMSG_MOVE_SET_SWIM_BACK_SPEED_CHEATMSG_MOVE_SET_SWIM_BACK_SPEEDMSG_MOVE_SET_ALL_SPEED_CHEATMSG_MOVE_SET_TURN_RATE_CHEATMSG_MOVE_SET_TURN_RATEMSG_MOVE_TOGGLE_COLLISION_CHEATMSG_MOVE_SET_FACINGMSG_MOVE_SET_PITCHMSG_MOVE_WORLDPORT_ACKSMSG_MONSTER_MOVESMSG_MOVE_WATER_WALKSMSG_MOVE_LAND_WALKCMSG_MOVE_CHARM_PORT_CHEATCMSG_MOVE_SET_RAW_POSITIONSMSG_FORCE_RUN_SPEED_CHANGECMSG_FORCE_RUN_SPEED_CHANGE_ACKSMSG_FORCE_RUN_BACK_SPEED_CHANGECMSG_FORCE_RUN_BACK_SPEED_CHANGE_ACKSMSG_FORCE_SWIM_SPEED_CHANGECMSG_FORCE_SWIM_SPEED_CHANGE_ACKSMSG_FORCE_MOVE_ROOTCMSG_FORCE_MOVE_ROOT_ACKSMSG_FORCE_MOVE_UNROOTCMSG_FORCE_MOVE_UNROOT_ACKMSG_MOVE_ROOTMSG_MOVE_UNROOTMSG_MOVE_HEARTBEATSMSG_MOVE_KNOCK_BACKCMSG_MOVE_KNOCK_BACK_ACKMSG_MOVE_KNOCK_BACKSMSG_MOVE_FEATHER_FALLSMSG_MOVE_NORMAL_FALLSMSG_MOVE_SET_HOVERSMSG_MOVE_UNSET_HOVERCMSG_MOVE_HOVER_ACKMSG_MOVE_HOVERCMSG_TRIGGER_CINEMATIC_CHEATCMSG_OPENING_CINEMATICSMSG_TRIGGER_CINEMATICCMSG_NEXT_CINEMATIC_CAMERACMSG_COMPLETE_CINEMATICSMSG_TUTORIAL_FLAGSCMSG_TUTORIAL_FLAGCMSG_TUTORIAL_CLEARCMSG_TUTORIAL_RESETCMSG_STANDSTATECHANGECMSG_EMOTESMSG_EMOTECMSG_TEXT_EMOTESMSG_TEXT_EMOTECMSG_AUTOEQUIP_GROUND_ITEMCMSG_AUTOSTORE_GROUND_ITEMCMSG_AUTOSTORE_LOOT_ITEMCMSG_STORE_LOOT_IN_SLOTCMSG_AUTOEQUIP_ITEMCMSG_AUTOSTORE_BAG_ITEMCMSG_SWAP_ITEMCMSG_SWAP_INV_ITEMCMSG_SPLIT_ITEMCMSG_AUTOEQUIP_ITEM_SLOTCMSG_UNCLAIM_LICENSECMSG_DESTROYITEMSMSG_INVENTORY_CHANGE_FAILURESMSG_OPEN_CONTAINERCMSG_INSPECTSMSG_INSPECT_RESULTS_UPDATECMSG_INITIATE_TRADECMSG_BEGIN_TRADECMSG_BUSY_TRADECMSG_IGNORE_TRADECMSG_ACCEPT_TRADECMSG_UNACCEPT_TRADECMSG_CANCEL_TRADECMSG_SET_TRADE_ITEMCMSG_CLEAR_TRADE_ITEMCMSG_SET_TRADE_GOLDSMSG_TRADE_STATUSSMSG_TRADE_STATUS_EXTENDEDSMSG_INITIALIZE_FACTIONSSMSG_SET_FACTION_VISIBLESMSG_SET_FACTION_STANDINGCMSG_SET_FACTION_ATWARCMSG_SET_FACTION_CHEATSMSG_SET_PROFICIENCYCMSG_SET_ACTION_BUTTONSMSG_ACTION_BUTTONSSMSG_INITIAL_SPELLSSMSG_LEARNED_SPELLSMSG_SUPERCEDED_SPELLCMSG_NEW_SPELL_SLOTCMSG_CAST_SPELLCMSG_CANCEL_CASTSMSG_CAST_FAILEDSMSG_SPELL_STARTSMSG_SPELL_GOSMSG_SPELL_FAILURESMSG_SPELL_COOLDOWNSMSG_COOLDOWN_EVENTCMSG_CANCEL_AURASMSG_EQUIPMENT_SET_SAVEDSMSG_PET_CAST_FAILEDMSG_CHANNEL_STARTMSG_CHANNEL_UPDATECMSG_CANCEL_CHANNELLINGSMSG_AI_REACTIONCMSG_SET_SELECTIONCMSG_DELETEEQUIPMENT_SETCMSG_INSTANCE_LOCK_RESPONSECMSG_DEBUG_PASSIVE_AURACMSG_ATTACKSWINGCMSG_ATTACKSTOPSMSG_ATTACKSTARTSMSG_ATTACKSTOPSMSG_ATTACKSWING_NOTINRANGESMSG_ATTACKSWING_BADFACINGSMSG_INSTANCE_LOCK_WARNING_QUERYSMSG_ATTACKSWING_DEADTARGETSMSG_ATTACKSWING_CANT_ATTACKSMSG_ATTACKERSTATEUPDATESMSG_BATTLEFIELD_PORT_DENIEDCMSG_PERFORM_ACTION_SETSMSG_RESUME_CAST_BARSMSG_CANCEL_COMBATSMSG_SPELLBREAKLOGSMSG_SPELLHEALLOGSMSG_SPELLENERGIZELOGSMSG_BREAK_TARGETCMSG_SAVE_PLAYERCMSG_SETDEATHBINDPOINTSMSG_BINDPOINTUPDATECMSG_GETDEATHBINDZONESMSG_BINDZONEREPLYSMSG_PLAYERBOUNDSMSG_CLIENT_CONTROL_UPDATECMSG_REPOP_REQUESTSMSG_RESURRECT_REQUESTCMSG_RESURRECT_RESPONSECMSG_LOOTCMSG_LOOT_MONEYCMSG_LOOT_RELEASESMSG_LOOT_RESPONSESMSG_LOOT_RELEASE_RESPONSESMSG_LOOT_REMOVEDSMSG_LOOT_MONEY_NOTIFYSMSG_LOOT_ITEM_NOTIFYSMSG_LOOT_CLEAR_MONEYSMSG_ITEM_PUSH_RESULTSMSG_DUEL_REQUESTEDSMSG_DUEL_OUTOFBOUNDSSMSG_DUEL_INBOUNDSSMSG_DUEL_COMPLETESMSG_DUEL_WINNERCMSG_DUEL_ACCEPTEDCMSG_DUEL_CANCELLEDSMSG_MOUNTRESULTSMSG_DISMOUNTRESULTSMSG_REMOVED_FROM_PVP_QUEUECMSG_MOUNTSPECIAL_ANIMSMSG_MOUNTSPECIAL_ANIMSMSG_PET_TAME_FAILURECMSG_PET_SET_ACTIONCMSG_PET_ACTIONCMSG_PET_ABANDONCMSG_PET_RENAMESMSG_PET_NAME_INVALIDSMSG_PET_SPELLSSMSG_PET_MODECMSG_GOSSIP_HELLOCMSG_GOSSIP_SELECT_OPTIONSMSG_GOSSIP_MESSAGESMSG_GOSSIP_COMPLETECMSG_NPC_TEXT_QUERYSMSG_NPC_TEXT_UPDATESMSG_NPC_WONT_TALKCMSG_QUESTGIVER_STATUS_QUERYSMSG_QUESTGIVER_STATUSCMSG_QUESTGIVER_HELLOSMSG_QUESTGIVER_QUEST_LISTCMSG_QUESTGIVER_QUERY_QUESTCMSG_QUESTGIVER_QUEST_AUTOLAUNCHSMSG_QUESTGIVER_QUEST_DETAILSCMSG_QUESTGIVER_ACCEPT_QUESTCMSG_QUESTGIVER_COMPLETE_QUESTSMSG_QUESTGIVER_REQUEST_ITEMSCMSG_QUESTGIVER_REQUEST_REWARDSMSG_QUESTGIVER_OFFER_REWARDCMSG_QUESTGIVER_CHOOSE_REWARDSMSG_QUESTGIVER_QUEST_INVALIDCMSG_QUESTGIVER_CANCELSMSG_QUESTGIVER_QUEST_COMPLETESMSG_QUESTGIVER_QUEST_FAILEDCMSG_QUESTLOG_SWAP_QUESTCMSG_QUESTLOG_REMOVE_QUESTSMSG_QUESTLOG_FULLSMSG_QUESTUPDATE_FAILEDSMSG_QUESTUPDATE_FAILEDTIMERSMSG_QUESTUPDATE_COMPLETESMSG_QUESTUPDATE_ADD_KILLSMSG_QUESTUPDATE_ADD_ITEMCMSG_QUEST_CONFIRM_ACCEPTSMSG_QUEST_CONFIRM_ACCEPTCMSG_PUSHQUESTTOPARTYCMSG_LIST_INVENTORYSMSG_LIST_INVENTORYCMSG_SELL_ITEMSMSG_SELL_ITEMCMSG_BUY_ITEMCMSG_BUY_ITEM_IN_SLOTSMSG_BUY_ITEMSMSG_BUY_FAILEDCMSG_TAXICLEARALLNODESCMSG_TAXIENABLEALLNODESCMSG_TAXISHOWNODESSMSG_SHOWTAXINODESCMSG_TAXINODE_STATUS_QUERYSMSG_TAXINODE_STATUSCMSG_TAXIQUERYAVAILABLENODESCMSG_ACTIVATETAXISMSG_ACTIVATETAXIREPLYSMSG_NEW_TAXI_PATHCMSG_TRAINER_LISTSMSG_TRAINER_LISTCMSG_TRAINER_BUY_SPELLSMSG_TRAINER_BUY_SUCCEEDEDSMSG_TRAINER_BUY_FAILEDCMSG_BINDER_ACTIVATESMSG_PLAYERBINDERRORCMSG_BANKER_ACTIVATESMSG_SHOW_BANKCMSG_BUY_BANK_SLOTSMSG_BUY_BANK_SLOT_RESULTCMSG_PETITION_SHOWLISTSMSG_PETITION_SHOWLISTCMSG_PETITION_BUYCMSG_PETITION_SHOW_SIGNATURESSMSG_PETITION_SHOW_SIGNATURESCMSG_PETITION_SIGNSMSG_PETITION_SIGN_RESULTSMSG_PETITION_DECLINECMSG_OFFER_PETITIONCMSG_TURN_IN_PETITIONSMSG_TURN_IN_PETITION_RESULTSCMSG_PETITION_QUERYSMSG_PETITION_QUERY_RESPONSESMSG_FISH_NOT_HOOKEDSMSG_FISH_ESCAPEDCMSG_BUGSMSG_NOTIFICATIONCMSG_PLAYED_TIMESMSG_PLAYED_TIMECMSG_QUERY_TIMESMSG_QUERY_TIME_RESPONSESMSG_LOG_XPGAINSMSG_AURACASTLOGCMSG_RECLAIM_CORPSECMSG_WRAP_ITEMSMSG_LEVELUP_INFOMSG_MINIMAP_PINGSMSG_RESISTLOGSMSG_ENCHANTMENTLOGCMSG_SET_SKILL_CHEATSMSG_START_MIRROR_TIMERSMSG_PAUSE_MIRROR_TIMERSMSG_STOP_MIRROR_TIMERCMSG_PINGSMSG_PONGSMSG_CLEAR_COOLDOWNSMSG_GAMEOBJECT_PAGETEXTCMSG_SETSHEATHEDSMSG_COOLDOWN_CHEATSMSG_SPELL_DELAYEDCMSG_QUEST_POI_QUERYSMSG_QUEST_POI_QUERY_RESPONSECMSG_GHOSTCMSG_GM_INVISSMSG_INVALID_PROMOTION_CODEMSG_GM_BIND_OTHERMSG_GM_SUMMONSMSG_ITEM_TIME_UPDATESMSG_ITEM_ENCHANT_TIME_UPDATESMSG_AUTH_CHALLENGECMSG_AUTH_SESSIONSMSG_AUTH_RESPONSEMSG_GM_SHOWLABELCMSG_PET_CAST_SPELLMSG_SAVE_GUILD_EMBLEMMSG_TABARDVENDOR_ACTIVATESMSG_PLAY_SPELL_VISUALCMSG_ZONEUPDATESMSG_PARTYKILLLOGSMSG_COMPRESSED_UPDATE_OBJECTSMSG_PLAY_SPELL_IMPACTSMSG_EXPLORATION_EXPERIENCECMSG_GM_SET_SECURITY_GROUPCMSG_GM_NUKEMSG_RANDOM_ROLLSMSG_ENVIRONMENTALDAMAGELOGCMSG_CHANGEPLAYER_DIFFICULTYSMSG_RWHOISSMSG_LFG_PLAYER_REWARDSMSG_LFG_TELEPORT_DENIEDCMSG_UNLEARN_SPELLCMSG_UNLEARN_SKILLSMSG_REMOVED_SPELLCMSG_DECHARGECMSG_GMTICKET_CREATESMSG_GMTICKET_CREATECMSG_GMTICKET_UPDATETEXTSMSG_GMTICKET_UPDATETEXTSMSG_ACCOUNT_DATA_TIMESCMSG_REQUEST_ACCOUNT_DATACMSG_UPDATE_ACCOUNT_DATASMSG_UPDATE_ACCOUNT_DATASMSG_CLEAR_FAR_SIGHT_IMMEDIATESMSG_CHANGEPLAYER_DIFFICULTY_RESULTCMSG_GM_TEACHCMSG_GM_CREATE_ITEM_TARGETCMSG_GMTICKET_GETTICKETSMSG_GMTICKET_GETTICKETCMSG_UNLEARN_TALENTSSMSG_UPDATE_INSTANCE_ENCOUNTER_UNITSMSG_GAMEOBJECT_DESPAWN_ANIMMSG_CORPSE_QUERYCMSG_GMTICKET_DELETETICKETSMSG_GMTICKET_DELETETICKETSMSG_CHAT_WRONG_FACTIONCMSG_GMTICKET_SYSTEMSTATUSSMSG_GMTICKET_SYSTEMSTATUSCMSG_SPIRIT_HEALER_ACTIVATECMSG_SET_STAT_CHEATSMSG_QUEST_FORCE_REMOVECMSG_SKILL_BUY_STEPCMSG_SKILL_BUY_RANKCMSG_XP_CHEATSMSG_SPIRIT_HEALER_CONFIRMCMSG_CHARACTER_POINT_CHEATSMSG_GOSSIP_POICMSG_CHAT_IGNOREDCMSG_GM_VISIONCMSG_SERVER_COMMANDCMSG_GM_SILENCECMSG_GM_REVEALTOCMSG_GM_RESURRECTCMSG_GM_SUMMONMOBCMSG_GM_MOVECORPSECMSG_GM_FREEZECMSG_GM_UBERINVISCMSG_GM_REQUEST_PLAYER_INFOSMSG_GM_PLAYER_INFOCMSG_GUILD_RANKCMSG_GUILD_ADD_RANKCMSG_GUILD_DEL_RANKCMSG_GUILD_SET_PUBLIC_NOTECMSG_GUILD_SET_OFFICER_NOTESMSG_LOGIN_VERIFY_WORLDCMSG_CLEAR_EXPLORATIONCMSG_SEND_MAILSMSG_SEND_MAIL_RESULTCMSG_GET_MAIL_LISTSMSG_MAIL_LIST_RESULTCMSG_BATTLEFIELD_LISTSMSG_BATTLEFIELD_LISTCMSG_BATTLEFIELD_JOINSMSG_FORCE_SET_VEHICLE_REC_IDCMSG_SET_VEHICLE_REC_ID_ACKCMSG_TAXICLEARNODECMSG_TAXIENABLENODECMSG_ITEM_TEXT_QUERYSMSG_ITEM_TEXT_QUERY_RESPONSECMSG_MAIL_TAKE_MONEYCMSG_MAIL_TAKE_ITEMCMSG_MAIL_MARK_AS_READCMSG_MAIL_RETURN_TO_SENDERCMSG_MAIL_DELETECMSG_MAIL_CREATE_TEXT_ITEMSMSG_SPELLLOGMISSSMSG_SPELLLOGEXECUTESMSG_DEBUGAURAPROCSMSG_PERIODICAURALOGSMSG_SPELLDAMAGESHIELDSMSG_SPELLNONMELEEDAMAGELOGCMSG_LEARN_TALENTSMSG_RESURRECT_FAILEDCMSG_TOGGLE_PVPSMSG_ZONE_UNDER_ATTACKMSG_AUCTION_HELLOCMSG_AUCTION_SELL_ITEMCMSG_AUCTION_REMOVE_ITEMCMSG_AUCTION_LIST_ITEMSCMSG_AUCTION_LIST_OWNER_ITEMSCMSG_AUCTION_PLACE_BIDSMSG_AUCTION_COMMAND_RESULTSMSG_AUCTION_LIST_RESULTSMSG_AUCTION_OWNER_LIST_RESULTSMSG_AUCTION_BIDDER_NOTIFICATIONSMSG_AUCTION_OWNER_NOTIFICATIONSMSG_PROCRESISTSMSG_COMBAT_EVENT_FAILEDSMSG_DISPEL_FAILEDSMSG_SPELLORDAMAGE_IMMUNECMSG_AUCTION_LIST_BIDDER_ITEMSSMSG_AUCTION_BIDDER_LIST_RESULTSMSG_SET_FLAT_SPELL_MODIFIERSMSG_SET_PCT_SPELL_MODIFIERCMSG_SET_AMMOSMSG_CORPSE_RECLAIM_DELAYCMSG_SET_ACTIVE_MOVERCMSG_PET_CANCEL_AURACMSG_PLAYER_AI_CHEATCMSG_CANCEL_AUTO_REPEAT_SPELLMSG_GM_ACCOUNT_ONLINEMSG_LIST_STABLED_PETSCMSG_STABLE_PETCMSG_UNSTABLE_PETCMSG_BUY_STABLE_SLOTSMSG_STABLE_RESULTCMSG_STABLE_REVIVE_PETCMSG_STABLE_SWAP_PETMSG_QUEST_PUSH_RESULTSMSG_PLAY_MUSICSMSG_PLAY_OBJECT_SOUNDCMSG_REQUEST_PET_INFOCMSG_FAR_SIGHTSMSG_SPELLDISPELLOGSMSG_DAMAGE_CALC_LOGCMSG_ENABLE_DAMAGE_LOGCMSG_GROUP_CHANGE_SUB_GROUPCMSG_REQUEST_PARTY_MEMBER_STATSCMSG_GROUP_SWAP_SUB_GROUPCMSG_RESET_FACTION_CHEATCMSG_AUTOSTORE_BANK_ITEMCMSG_AUTOBANK_ITEMMSG_QUERY_NEXT_MAIL_TIMESMSG_RECEIVED_MAILSMSG_RAID_GROUP_ONLYCMSG_SET_DURABILITY_CHEATCMSG_SET_PVP_RANK_CHEATCMSG_ADD_PVP_MEDAL_CHEATCMSG_DEL_PVP_MEDAL_CHEATCMSG_SET_PVP_TITLESMSG_PVP_CREDITSMSG_AUCTION_REMOVED_NOTIFICATIONCMSG_GROUP_RAID_CONVERTCMSG_GROUP_ASSISTANT_LEADERCMSG_BUYBACK_ITEMSMSG_SERVER_MESSAGECMSG_SET_SAVED_INSTANCE_EXTENDSMSG_LFG_OFFER_CONTINUECMSG_TEST_DROP_RATESMSG_TEST_DROP_RATE_RESULTCMSG_LFG_GET_STATUSSMSG_SHOW_MAILBOXSMSG_RESET_RANGED_COMBAT_TIMERSMSG_CHAT_NOT_IN_PARTYCMSG_GMTICKETSYSTEM_TOGGLECMSG_CANCEL_GROWTH_AURASMSG_CANCEL_AUTO_REPEATSMSG_STANDSTATE_UPDATESMSG_LOOT_ALL_PASSEDSMSG_LOOT_ROLL_WONCMSG_LOOT_ROLLSMSG_LOOT_START_ROLLSMSG_LOOT_ROLLCMSG_LOOT_MASTER_GIVESMSG_LOOT_MASTER_LISTSMSG_SET_FORCED_REACTIONSSMSG_SPELL_FAILED_OTHERSMSG_GAMEOBJECT_RESET_STATECMSG_REPAIR_ITEMSMSG_CHAT_PLAYER_NOT_FOUNDMSG_TALENT_WIPE_CONFIRMSMSG_SUMMON_REQUESTCMSG_SUMMON_RESPONSEMSG_DEV_SHOWLABELSMSG_MONSTER_MOVE_TRANSPORTSMSG_PET_BROKENMSG_MOVE_FEATHER_FALLMSG_MOVE_WATER_WALKCMSG_SERVER_BROADCASTCMSG_SELF_RESSMSG_FEIGN_DEATH_RESISTEDCMSG_RUN_SCRIPTSMSG_SCRIPT_MESSAGESMSG_DUEL_COUNTDOWNSMSG_AREA_TRIGGER_MESSAGECMSG_SHOWING_HELMCMSG_SHOWING_CLOAKSMSG_LFG_ROLE_CHOSENSMSG_PLAYER_SKINNEDSMSG_DURABILITY_DAMAGE_DEATHCMSG_SET_EXPLORATIONCMSG_SET_ACTIONBAR_TOGGLESUMSG_DELETE_GUILD_CHARTERMSG_PETITION_RENAMESMSG_INIT_WORLD_STATESSMSG_UPDATE_WORLD_STATECMSG_ITEM_NAME_QUERYSMSG_ITEM_NAME_QUERY_RESPONSESMSG_PET_ACTION_FEEDBACKCMSG_CHAR_RENAMESMSG_CHAR_RENAMECMSG_MOVE_SPLINE_DONECMSG_MOVE_FALL_RESETSMSG_INSTANCE_SAVE_CREATEDSMSG_RAID_INSTANCE_INFOCMSG_REQUEST_RAID_INFOCMSG_MOVE_TIME_SKIPPEDCMSG_MOVE_FEATHER_FALL_ACKCMSG_MOVE_WATER_WALK_ACKCMSG_MOVE_NOT_ACTIVE_MOVERSMSG_PLAY_SOUNDCMSG_BATTLEFIELD_STATUSSMSG_BATTLEFIELD_STATUSCMSG_BATTLEFIELD_PORTMSG_INSPECT_HONOR_STATSCMSG_BATTLEMASTER_HELLOCMSG_MOVE_START_SWIM_CHEATCMSG_MOVE_STOP_SWIM_CHEATSMSG_FORCE_WALK_SPEED_CHANGECMSG_FORCE_WALK_SPEED_CHANGE_ACKSMSG_FORCE_SWIM_BACK_SPEED_CHANGECMSG_FORCE_SWIM_BACK_SPEED_CHANGE_ACKSMSG_FORCE_TURN_RATE_CHANGECMSG_FORCE_TURN_RATE_CHANGE_ACKMSG_PVP_LOG_DATACMSG_LEAVE_BATTLEFIELDCMSG_AREA_SPIRIT_HEALER_QUERYCMSG_AREA_SPIRIT_HEALER_QUEUESMSG_AREA_SPIRIT_HEALER_TIMECMSG_GM_UNTEACHSMSG_WARDEN_DATACMSG_WARDEN_DATASMSG_GROUP_JOINED_BATTLEGROUNDMSG_BATTLEGROUND_PLAYER_POSITIONSCMSG_PET_STOP_ATTACKSMSG_BINDER_CONFIRMSMSG_BATTLEGROUND_PLAYER_JOINEDSMSG_BATTLEGROUND_PLAYER_LEFTCMSG_BATTLEMASTER_JOINSMSG_ADDON_INFOCMSG_PET_UNLEARNSMSG_PET_UNLEARN_CONFIRMSMSG_PARTY_MEMBER_STATS_FULLCMSG_PET_SPELL_AUTOCASTSMSG_WEATHERSMSG_PLAY_TIME_WARNINGSMSG_MINIGAME_SETUPSMSG_MINIGAME_STATECMSG_MINIGAME_MOVESMSG_MINIGAME_MOVE_FAILEDSMSG_RAID_INSTANCE_MESSAGESMSG_COMPRESSED_MOVESCMSG_GUILD_INFO_TEXTSMSG_CHAT_RESTRICTEDSMSG_SPLINE_SET_RUN_SPEEDSMSG_SPLINE_SET_RUN_BACK_SPEEDSMSG_SPLINE_SET_SWIM_SPEEDSMSG_SPLINE_SET_WALK_SPEEDSMSG_SPLINE_SET_SWIM_BACK_SPEEDSMSG_SPLINE_SET_TURN_RATESMSG_SPLINE_MOVE_UNROOTSMSG_SPLINE_MOVE_FEATHER_FALLSMSG_SPLINE_MOVE_NORMAL_FALLSMSG_SPLINE_MOVE_SET_HOVERSMSG_SPLINE_MOVE_UNSET_HOVERSMSG_SPLINE_MOVE_WATER_WALKSMSG_SPLINE_MOVE_LAND_WALKSMSG_SPLINE_MOVE_START_SWIMSMSG_SPLINE_MOVE_STOP_SWIMSMSG_SPLINE_MOVE_SET_RUN_MODESMSG_SPLINE_MOVE_SET_WALK_MODECMSG_GM_NUKE_ACCOUNTMSG_GM_DESTROY_CORPSECMSG_GM_DESTROY_ONLINE_CORPSECMSG_ACTIVATETAXIEXPRESSSMSG_SET_FACTION_ATWARSMSG_GAMETIMEBIAS_SETCMSG_DEBUG_ACTIONS_STARTCMSG_DEBUG_ACTIONS_STOPCMSG_SET_FACTION_INACTIVECMSG_SET_WATCHED_FACTIONMSG_MOVE_TIME_SKIPPEDSMSG_SPLINE_MOVE_ROOTCMSG_SET_EXPLORATION_ALLSMSG_INVALIDATE_PLAYERCMSG_RESET_INSTANCESSMSG_INSTANCE_RESETSMSG_INSTANCE_RESET_FAILEDSMSG_UPDATE_LAST_INSTANCEMSG_RAID_TARGET_UPDATEMSG_RAID_READY_CHECKCMSG_LUA_USAGESMSG_PET_ACTION_SOUNDSMSG_PET_DISMISS_SOUNDSMSG_GHOSTEE_GONECMSG_GM_UPDATE_TICKET_STATUSSMSG_GM_TICKET_STATUS_UPDATEMSG_SET_DUNGEON_DIFFICULTYCMSG_GMSURVEY_SUBMITSMSG_UPDATE_INSTANCE_OWNERSHIPCMSG_IGNORE_KNOCKBACK_CHEATSMSG_CHAT_PLAYER_AMBIGUOUSMSG_DELAY_GHOST_TELEPORTSMSG_SPELLINSTAKILLLOGSMSG_SPELL_UPDATE_CHAIN_TARGETSCMSG_CHAT_FILTEREDSMSG_EXPECTED_SPAM_RECORDSSMSG_SPELLSTEALLOGCMSG_LOTTERY_QUERY_OBSOLETESMSG_LOTTERY_QUERY_RESULT_OBSOLETECMSG_BUY_LOTTERY_TICKET_OBSOLETESMSG_LOTTERY_RESULT_OBSOLETESMSG_CHARACTER_PROFILESMSG_CHARACTER_PROFILE_REALM_CONNECTEDSMSG_DEFENSE_MESSAGESMSG_INSTANCE_DIFFICULTYMSG_GM_RESETINSTANCELIMITSMSG_MOTDSMSG_MOVE_SET_CAN_TRANSITION_BETWEEN_SWIM_AND_FLYSMSG_MOVE_UNSET_CAN_TRANSITION_BETWEEN_SWIM_AND_FLYCMSG_MOVE_SET_CAN_TRANSITION_BETWEEN_SWIM_AND_FLY_ACKMSG_MOVE_START_SWIM_CHEATMSG_MOVE_STOP_SWIM_CHEATSMSG_MOVE_SET_CAN_FLYSMSG_MOVE_UNSET_CAN_FLYCMSG_MOVE_SET_CAN_FLY_ACKCMSG_MOVE_SET_FLYCMSG_SOCKET_GEMSCMSG_ARENA_TEAM_CREATESMSG_ARENA_TEAM_COMMAND_RESULTMSG_MOVE_UPDATE_CAN_TRANSITION_BETWEEN_SWIM_AND_FLYCMSG_ARENA_TEAM_QUERYSMSG_ARENA_TEAM_QUERY_RESPONSECMSG_ARENA_TEAM_ROSTERSMSG_ARENA_TEAM_ROSTERCMSG_ARENA_TEAM_INVITESMSG_ARENA_TEAM_INVITECMSG_ARENA_TEAM_ACCEPTCMSG_ARENA_TEAM_DECLINECMSG_ARENA_TEAM_LEAVECMSG_ARENA_TEAM_REMOVECMSG_ARENA_TEAM_DISBANDCMSG_ARENA_TEAM_LEADERSMSG_ARENA_TEAM_EVENTCMSG_BATTLEMASTER_JOIN_ARENAMSG_MOVE_START_ASCENDMSG_MOVE_STOP_ASCENDSMSG_ARENA_TEAM_STATSCMSG_LFG_JOINCMSG_LFG_LEAVECMSG_SEARCH_LFG_JOINCMSG_SEARCH_LFG_LEAVESMSG_UPDATE_LFG_LISTSMSG_LFG_PROPOSAL_UPDATECMSG_LFG_PROPOSAL_RESULTSMSG_LFG_ROLE_CHECK_UPDATESMSG_LFG_JOIN_RESULTSMSG_LFG_QUEUE_STATUSCMSG_SET_LFG_COMMENTSMSG_LFG_UPDATE_PLAYERSMSG_LFG_UPDATE_PARTYSMSG_LFG_UPDATE_SEARCHCMSG_LFG_SET_ROLESCMSG_LFG_SET_NEEDSCMSG_LFG_SET_BOOT_VOTESMSG_LFG_BOOT_PROPOSAL_UPDATECMSG_LFD_PLAYER_LOCK_INFO_REQUESTSMSG_LFG_PLAYER_INFOCMSG_LFG_TELEPORTCMSG_LFD_PARTY_LOCK_INFO_REQUESTSMSG_LFG_PARTY_INFOSMSG_TITLE_EARNEDCMSG_SET_TITLECMSG_CANCEL_MOUNT_AURASMSG_ARENA_ERRORMSG_INSPECT_ARENA_TEAMSSMSG_DEATH_RELEASE_LOCCMSG_CANCEL_TEMP_ENCHANTMENTSMSG_FORCED_DEATH_UPDATECMSG_CHEAT_SET_HONOR_CURRENCYCMSG_CHEAT_SET_ARENA_CURRENCYMSG_MOVE_SET_FLIGHT_SPEED_CHEATMSG_MOVE_SET_FLIGHT_SPEEDMSG_MOVE_SET_FLIGHT_BACK_SPEED_CHEATMSG_MOVE_SET_FLIGHT_BACK_SPEEDSMSG_FORCE_FLIGHT_SPEED_CHANGECMSG_FORCE_FLIGHT_SPEED_CHANGE_ACKSMSG_FORCE_FLIGHT_BACK_SPEED_CHANGECMSG_FORCE_FLIGHT_BACK_SPEED_CHANGE_ACKSMSG_SPLINE_SET_FLIGHT_SPEEDSMSG_SPLINE_SET_FLIGHT_BACK_SPEEDCMSG_MAELSTROM_INVALIDATE_CACHESMSG_FLIGHT_SPLINE_SYNCCMSG_SET_TAXI_BENCHMARK_MODESMSG_JOINED_BATTLEGROUND_QUEUESMSG_REALM_SPLITCMSG_REALM_SPLITCMSG_MOVE_CHNG_TRANSPORTMSG_PARTY_ASSIGNMENTSMSG_OFFER_PETITION_ERRORSMSG_TIME_SYNC_REQCMSG_TIME_SYNC_RESPCMSG_SEND_LOCAL_EVENTCMSG_SEND_GENERAL_TRIGGERCMSG_SEND_COMBAT_TRIGGERCMSG_MAELSTROM_GM_SENT_MAILSMSG_RESET_FAILED_NOTIFYSMSG_REAL_GROUP_UPDATESMSG_LFG_DISABLEDCMSG_ACTIVE_PVP_CHEATCMSG_CHEAT_DUMP_ITEMS_DEBUG_ONLYSMSG_CHEAT_DUMP_ITEMS_DEBUG_ONLY_RESPONSESMSG_CHEAT_DUMP_ITEMS_DEBUG_ONLY_RESPONSE_WRITE_FILESMSG_UPDATE_COMBO_POINTSSMSG_VOICE_SESSION_ROSTER_UPDATESMSG_VOICE_SESSION_LEAVESMSG_VOICE_SESSION_ADJUST_PRIORITYCMSG_VOICE_SET_TALKER_MUTED_REQUESTSMSG_VOICE_SET_TALKER_MUTEDSMSG_INIT_EXTRA_AURA_INFO_OBSOLETESMSG_SET_EXTRA_AURA_INFO_OBSOLETESMSG_SET_EXTRA_AURA_INFO_NEED_UPDATE_OBSOLETESMSG_CLEAR_EXTRA_AURA_INFO_OBSOLETEMSG_MOVE_START_DESCENDCMSG_IGNORE_REQUIREMENTS_CHEATSMSG_IGNORE_REQUIREMENTS_CHEATSMSG_SPELL_CHANCE_PROC_LOGCMSG_MOVE_SET_RUN_SPEEDSMSG_DISMOUNTMSG_MOVE_UPDATE_CAN_FLYMSG_RAID_READY_CHECK_CONFIRMCMSG_VOICE_SESSION_ENABLESMSG_VOICE_SESSION_ENABLESMSG_VOICE_PARENTAL_CONTROLSCMSG_GM_WHISPERSMSG_GM_MESSAGECHATMSG_GM_GEARRATINGCMSG_COMMENTATOR_ENABLESMSG_COMMENTATOR_STATE_CHANGEDCMSG_COMMENTATOR_GET_MAP_INFOSMSG_COMMENTATOR_MAP_INFOCMSG_COMMENTATOR_GET_PLAYER_INFOSMSG_COMMENTATOR_GET_PLAYER_INFOSMSG_COMMENTATOR_PLAYER_INFOCMSG_COMMENTATOR_ENTER_INSTANCECMSG_COMMENTATOR_EXIT_INSTANCECMSG_COMMENTATOR_INSTANCE_COMMANDSMSG_CLEAR_TARGETCMSG_BOT_DETECTEDSMSG_CROSSED_INEBRIATION_THRESHOLDCMSG_CHEAT_PLAYER_LOGINCMSG_CHEAT_PLAYER_LOOKUPSMSG_CHEAT_PLAYER_LOOKUPSMSG_KICK_REASONMSG_RAID_READY_CHECK_FINISHEDCMSG_COMPLAINSMSG_COMPLAIN_RESULTSMSG_FEATURE_SYSTEM_STATUSCMSG_GM_SHOW_COMPLAINTSCMSG_GM_UNSQUELCHCMSG_CHANNEL_SILENCE_VOICECMSG_CHANNEL_SILENCE_ALLCMSG_CHANNEL_UNSILENCE_VOICECMSG_CHANNEL_UNSILENCE_ALLCMSG_TARGET_CASTCMSG_TARGET_SCRIPT_CASTCMSG_CHANNEL_DISPLAY_LISTCMSG_SET_ACTIVE_VOICE_CHANNELCMSG_GET_CHANNEL_MEMBER_COUNTSMSG_CHANNEL_MEMBER_COUNTCMSG_CHANNEL_VOICE_ONCMSG_CHANNEL_VOICE_OFFCMSG_DEBUG_LIST_TARGETSSMSG_DEBUG_LIST_TARGETSSMSG_AVAILABLE_VOICE_CHANNELCMSG_ADD_VOICE_IGNORECMSG_DEL_VOICE_IGNORECMSG_PARTY_SILENCECMSG_PARTY_UNSILENCEMSG_NOTIFY_PARTY_SQUELCHSMSG_COMSAT_RECONNECT_TRYSMSG_COMSAT_DISCONNECTSMSG_COMSAT_CONNECT_FAILSMSG_VOICE_CHAT_STATUSCMSG_REPORT_PVP_AFKSMSG_REPORT_PVP_AFK_RESULTCMSG_GUILD_BANKER_ACTIVATECMSG_GUILD_BANK_QUERY_TABSMSG_GUILD_BANK_LISTCMSG_GUILD_BANK_SWAP_ITEMSCMSG_GUILD_BANK_BUY_TABCMSG_GUILD_BANK_UPDATE_TABCMSG_GUILD_BANK_DEPOSIT_MONEYCMSG_GUILD_BANK_WITHDRAW_MONEYMSG_GUILD_BANK_LOG_QUERYCMSG_SET_CHANNEL_WATCHSMSG_USERLIST_ADDSMSG_USERLIST_REMOVESMSG_USERLIST_UPDATECMSG_CLEAR_CHANNEL_WATCHSMSG_INSPECT_TALENTSMSG_GOGOGO_OBSOLETESMSG_ECHO_PARTY_SQUELCHCMSG_SET_TITLE_SUFFIXCMSG_SPELLCLICKSMSG_LOOT_LISTCMSG_GM_CHARACTER_RESTORECMSG_GM_CHARACTER_SAVESMSG_VOICESESSION_FULLMSG_GUILD_PERMISSIONSMSG_GUILD_BANK_MONEY_WITHDRAWNMSG_GUILD_EVENT_LOG_QUERYCMSG_MAELSTROM_RENAME_GUILDCMSG_GET_MIRRORIMAGE_DATASMSG_MIRRORIMAGE_DATASMSG_FORCE_DISPLAY_UPDATESMSG_SPELL_CHANCE_RESIST_PUSHBACKCMSG_IGNORE_DIMINISHING_RETURNS_CHEATSMSG_IGNORE_DIMINISHING_RETURNS_CHEATCMSG_KEEP_ALIVESMSG_RAID_READY_CHECK_ERRORCMSG_OPT_OUT_OF_LOOTMSG_QUERY_GUILD_BANK_TEXTCMSG_SET_GUILD_BANK_TEXTCMSG_SET_GRANTABLE_LEVELSCMSG_GRANT_LEVELCMSG_REFER_A_FRIENDMSG_GM_CHANGE_ARENA_RATINGCMSG_DECLINE_CHANNEL_INVITESMSG_GROUPACTION_THROTTLEDSMSG_OVERRIDE_LIGHTSMSG_TOTEM_CREATEDCMSG_TOTEM_DESTROYEDCMSG_EXPIRE_RAID_INSTANCECMSG_NO_SPELL_VARIANCECMSG_QUESTGIVER_STATUS_MULTIPLE_QUERYSMSG_QUESTGIVER_STATUS_MULTIPLECMSG_SET_PLAYER_DECLINED_NAMESSMSG_SET_PLAYER_DECLINED_NAMES_RESULTCMSG_QUERY_SERVER_BUCK_DATACMSG_CLEAR_SERVER_BUCK_DATASMSG_SERVER_BUCK_DATASMSG_SEND_UNLEARN_SPELLSSMSG_PROPOSE_LEVEL_GRANTCMSG_ACCEPT_LEVEL_GRANTSMSG_REFER_A_FRIEND_FAILURESMSG_SPLINE_MOVE_SET_FLYINGSMSG_SPLINE_MOVE_UNSET_FLYINGSMSG_SUMMON_CANCELCMSG_CHANGE_PERSONAL_ARENA_RATINGCMSG_ALTER_APPEARANCESMSG_ENABLE_BARBER_SHOPSMSG_BARBER_SHOP_RESULTCMSG_CALENDAR_GET_CALENDARCMSG_CALENDAR_GET_EVENTCMSG_CALENDAR_GUILD_FILTERCMSG_CALENDAR_ARENA_TEAMCMSG_CALENDAR_ADD_EVENTCMSG_CALENDAR_UPDATE_EVENTCMSG_CALENDAR_REMOVE_EVENTCMSG_CALENDAR_COPY_EVENTCMSG_CALENDAR_EVENT_INVITECMSG_CALENDAR_EVENT_RSVPCMSG_CALENDAR_EVENT_REMOVE_INVITECMSG_CALENDAR_EVENT_STATUSCMSG_CALENDAR_EVENT_MODERATOR_STATUSSMSG_CALENDAR_SEND_CALENDARSMSG_CALENDAR_SEND_EVENTSMSG_CALENDAR_FILTER_GUILDSMSG_CALENDAR_ARENA_TEAMSMSG_CALENDAR_EVENT_INVITESMSG_CALENDAR_EVENT_INVITE_REMOVEDSMSG_CALENDAR_EVENT_STATUSSMSG_CALENDAR_COMMAND_RESULTSMSG_CALENDAR_RAID_LOCKOUT_ADDEDSMSG_CALENDAR_RAID_LOCKOUT_REMOVEDSMSG_CALENDAR_EVENT_INVITE_ALERTSMSG_CALENDAR_EVENT_INVITE_REMOVED_ALERTSMSG_CALENDAR_EVENT_INVITE_STATUS_ALERTSMSG_CALENDAR_EVENT_REMOVED_ALERTSMSG_CALENDAR_EVENT_UPDATED_ALERTSMSG_CALENDAR_EVENT_MODERATOR_STATUS_ALERTCMSG_CALENDAR_COMPLAINCMSG_CALENDAR_GET_NUM_PENDINGSMSG_CALENDAR_SEND_NUM_PENDINGCMSG_SAVE_DANCESMSG_NOTIFY_DANCECMSG_PLAY_DANCESMSG_PLAY_DANCECMSG_LOAD_DANCESCMSG_STOP_DANCESMSG_STOP_DANCECMSG_SYNC_DANCECMSG_DANCE_QUERYSMSG_DANCE_QUERY_RESPONSESMSG_INVALIDATE_DANCECMSG_DELETE_DANCESMSG_LEARNED_DANCE_MOVESCMSG_LEARN_DANCE_MOVECMSG_UNLEARN_DANCE_MOVECMSG_SET_RUNE_COUNTCMSG_SET_RUNE_COOLDOWNMSG_MOVE_SET_PITCH_RATE_CHEATMSG_MOVE_SET_PITCH_RATESMSG_FORCE_PITCH_RATE_CHANGECMSG_FORCE_PITCH_RATE_CHANGE_ACKSMSG_SPLINE_SET_PITCH_RATECMSG_CALENDAR_EVENT_INVITE_NOTESSMSG_CALENDAR_EVENT_INVITE_NOTESSMSG_CALENDAR_EVENT_INVITE_NOTES_ALERTCMSG_UPDATE_MISSILE_TRAJECTORYSMSG_UPDATE_ACCOUNT_DATA_COMPLETESMSG_TRIGGER_MOVIECMSG_COMPLETE_MOVIECMSG_SET_GLYPH_SLOTCMSG_SET_GLYPHSMSG_ACHIEVEMENT_EARNEDSMSG_DYNAMIC_DROP_ROLL_RESULTSMSG_CRITERIA_UPDATECMSG_QUERY_INSPECT_ACHIEVEMENTSSMSG_RESPOND_INSPECT_ACHIEVEMENTSCMSG_DISMISS_CONTROLLED_VEHICLECMSG_COMPLETE_ACHIEVEMENT_CHEATSMSG_QUESTUPDATE_ADD_PVP_KILLCMSG_SET_CRITERIA_CHEATSMSG_CALENDAR_RAID_LOCKOUT_UPDATEDCMSG_UNITANIMTIER_CHEATCMSG_CHAR_CUSTOMIZESMSG_CHAR_CUSTOMIZESMSG_PET_RENAMEABLECMSG_REQUEST_VEHICLE_EXITCMSG_REQUEST_VEHICLE_PREV_SEATCMSG_REQUEST_VEHICLE_NEXT_SEATCMSG_REQUEST_VEHICLE_SWITCH_SEATCMSG_PET_LEARN_TALENTCMSG_PET_UNLEARN_TALENTSSMSG_SET_PHASE_SHIFTSMSG_ALL_ACHIEVEMENT_DATACMSG_FORCE_SAY_CHEATSMSG_HEALTH_UPDATESMSG_POWER_UPDATECMSG_GAMEOBJ_REPORT_USESMSG_HIGHEST_THREAT_UPDATESMSG_THREAT_UPDATESMSG_THREAT_REMOVESMSG_THREAT_CLEARSMSG_CONVERT_RUNESMSG_RESYNC_RUNESSMSG_ADD_RUNE_POWERCMSG_START_QUESTCMSG_REMOVE_GLYPHCMSG_DUMP_OBJECTSSMSG_DUMP_OBJECTS_DATACMSG_DISMISS_CRITTERSMSG_NOTIFY_DEST_LOC_SPELL_CASTCMSG_AUCTION_LIST_PENDING_SALESSMSG_AUCTION_LIST_PENDING_SALESSMSG_MODIFY_COOLDOWNSMSG_PET_UPDATE_COMBO_POINTSCMSG_ENABLETAXISMSG_PRE_RESURRECTSMSG_AURA_UPDATE_ALLSMSG_AURA_UPDATECMSG_FLOOD_GRACE_CHEATSMSG_SERVER_FIRST_ACHIEVEMENTSMSG_PET_LEARNED_SPELLSMSG_PET_REMOVED_SPELLCMSG_CHANGE_SEATS_ON_CONTROLLED_VEHICLECMSG_HEARTH_AND_RESURRECTSMSG_ON_CANCEL_EXPECTED_RIDE_VEHICLE_AURASMSG_CRITERIA_DELETEDSMSG_ACHIEVEMENT_DELETEDCMSG_SERVER_INFO_QUERYSMSG_SERVER_INFO_RESPONSECMSG_CHECK_LOGIN_CRITERIASMSG_SERVER_BUCK_DATA_STARTCMSG_SET_BREATHCMSG_QUERY_VEHICLE_STATUSSMSG_BATTLEGROUND_INFO_THROTTLEDSMSG_PLAYER_VEHICLE_DATACMSG_PLAYER_VEHICLE_ENTERCMSG_CONTROLLER_EJECT_PASSENGERSMSG_PET_GUIDSSMSG_CLIENTCACHE_VERSIONCMSG_CHANGE_GDF_ARENA_RATINGCMSG_SET_ARENA_TEAM_RATING_BY_INDEXCMSG_SET_ARENA_TEAM_WEEKLY_GAMESCMSG_SET_ARENA_TEAM_SEASON_GAMESCMSG_SET_ARENA_MEMBER_WEEKLY_GAMESCMSG_SET_ARENA_MEMBER_SEASON_GAMESSMSG_ITEM_REFUND_INFO_RESPONSECMSG_ITEM_REFUND_INFOCMSG_ITEM_REFUNDSMSG_ITEM_REFUND_RESULTCMSG_CORPSE_MAP_POSITION_QUERYSMSG_CORPSE_MAP_POSITION_QUERY_RESPONSECMSG_UNUSED5CMSG_UNUSED6CMSG_CALENDAR_EVENT_SIGNUPSMSG_CALENDAR_CLEAR_PENDING_ACTIONSMSG_EQUIPMENT_SET_LISTCMSG_EQUIPMENT_SET_SAVECMSG_UPDATE_PROJECTILE_POSITIONSMSG_SET_PROJECTILE_POSITIONSMSG_TALENTS_INFOCMSG_LEARN_PREVIEW_TALENTSCMSG_LEARN_PREVIEW_TALENTS_PETCMSG_SET_ACTIVE_TALENT_GROUP_OBSOLETECMSG_GM_GRANT_ACHIEVEMENTCMSG_GM_REMOVE_ACHIEVEMENTCMSG_GM_SET_CRITERIA_FOR_PLAYERSMSG_ARENA_UNIT_DESTROYEDSMSG_ARENA_TEAM_CHANGE_FAILED_QUEUEDCMSG_PROFILEDATA_REQUESTSMSG_PROFILEDATA_RESPONSECMSG_START_BATTLEFIELD_CHEATCMSG_END_BATTLEFIELD_CHEATSMSG_MULTIPLE_PACKETSSMSG_MOVE_GRAVITY_DISABLECMSG_MOVE_GRAVITY_DISABLE_ACKSMSG_MOVE_GRAVITY_ENABLECMSG_MOVE_GRAVITY_ENABLE_ACKMSG_MOVE_GRAVITY_CHNGSMSG_SPLINE_MOVE_GRAVITY_DISABLESMSG_SPLINE_MOVE_GRAVITY_ENABLECMSG_EQUIPMENT_SET_USESMSG_EQUIPMENT_SET_USE_RESULTCMSG_FORCE_ANIMSMSG_FORCE_ANIMCMSG_CHAR_FACTION_CHANGESMSG_CHAR_FACTION_CHANGECMSG_PVP_QUEUE_STATS_REQUESTSMSG_PVP_QUEUE_STATSCMSG_SET_PAID_SERVICE_CHEATSMSG_BATTLEFIELD_MGR_ENTRY_INVITECMSG_BATTLEFIELD_MGR_ENTRY_INVITE_RESPONSESMSG_BATTLEFIELD_MGR_ENTEREDSMSG_BATTLEFIELD_MGR_QUEUE_INVITECMSG_BATTLEFIELD_MGR_QUEUE_INVITE_RESPONSECMSG_BATTLEFIELD_MGR_QUEUE_REQUESTSMSG_BATTLEFIELD_MGR_QUEUE_REQUEST_RESPONSESMSG_BATTLEFIELD_MGR_EJECT_PENDINGSMSG_BATTLEFIELD_MGR_EJECTEDCMSG_BATTLEFIELD_MGR_EXIT_REQUESTSMSG_BATTLEFIELD_MGR_STATE_CHANGECMSG_BATTLEFIELD_MANAGER_ADVANCE_STATECMSG_BATTLEFIELD_MANAGER_SET_NEXT_TRANSITION_TIMEMSG_SET_RAID_DIFFICULTYCMSG_TOGGLE_XP_GAINSMSG_TOGGLE_XP_GAINSMSG_GMRESPONSE_DB_ERRORSMSG_GMRESPONSE_RECEIVEDCMSG_GMRESPONSE_RESOLVESMSG_GMRESPONSE_STATUS_UPDATESMSG_GMRESPONSE_CREATE_TICKETCMSG_GMRESPONSE_CREATE_TICKETCMSG_SERVERINFOSMSG_SERVERINFOCMSG_WORLD_STATE_UI_TIMER_UPDATESMSG_WORLD_STATE_UI_TIMER_UPDATECMSG_CHAR_RACE_CHANGEMSG_VIEW_PHASE_SHIFTSMSG_TALENTS_INVOLUNTARILY_RESETCMSG_DEBUG_SERVER_GEOSMSG_DEBUG_SERVER_GEOSMSG_LOOT_SLOT_CHANGEDUMSG_UPDATE_GROUP_INFOCMSG_READY_FOR_ACCOUNT_DATA_TIMESCMSG_QUERY_QUESTS_COMPLETEDSMSG_QUERY_QUESTS_COMPLETED_RESPONSECMSG_GM_REPORT_LAGCMSG_AFK_MONITOR_INFO_REQUESTSMSG_AFK_MONITOR_INFO_RESPONSECMSG_AFK_MONITOR_INFO_CLEARSMSG_CORPSE_NOT_IN_INSTANCECMSG_GM_NUKE_CHARACTERCMSG_SET_ALLOW_LOW_LEVEL_RAID1CMSG_SET_ALLOW_LOW_LEVEL_RAID2SMSG_CAMERA_SHAKESMSG_SOCKET_GEMS_RESULTCMSG_SET_CHARACTER_MODELSMSG_REDIRECT_CLIENTCMSG_REDIRECTION_FAILEDSMSG_SUSPEND_COMMSCMSG_SUSPEND_COMMS_ACKSMSG_FORCE_SEND_QUEUED_PACKETSCMSG_REDIRECTION_AUTH_PROOFCMSG_DROP_NEW_CONNECTIONSMSG_SEND_ALL_COMBAT_LOGSMSG_OPEN_LFG_DUNGEON_FINDERSMSG_MOVE_SET_COLLISION_HGTCMSG_MOVE_SET_COLLISION_HGT_ACKMSG_MOVE_SET_COLLISION_HGTCMSG_CLEAR_RANDOM_BG_WIN_TIMECMSG_CLEAR_HOLIDAY_BG_WIN_TIMECMSG_COMMENTATOR_SKIRMISH_QUEUE_COMMANDSMSG_COMMENTATOR_SKIRMISH_QUEUE_RESULT1SMSG_COMMENTATOR_SKIRMISH_QUEUE_RESULT2SMSG_MULTIPLE_MOVESNUM_MSG_TYPES"

var _WorldType_index = [...]uint16{0, 11, 24, 37, 63, 89, 115, 141, 160, 181, 194, 207, 232, 257, 282, 295, 311, 329, 348, 363, 384, 403, 429, 447, 463, 486, 506, 526, 541, 556, 583, 607, 625, 641, 653, 665, 684, 700, 720, 739, 758, 778, 793, 815, 831, 846, 864, 882, 904, 927, 960, 980, 1000, 1021, 1037, 1051, 1067, 1090, 1106, 1120, 1136, 1153, 1167, 1188, 1209, 1236, 1259, 1279, 1296, 1313, 1331, 1349, 1364, 1379, 1397, 1416, 1436, 1456, 1474, 1496, 1511, 1535, 1554, 1582, 1598, 1623, 1645, 1669, 1700, 1733, 1753, 1782, 1798, 1823, 1844, 1874, 1893, 1921, 1929, 1937, 1947, 1957, 1974, 1991, 2009, 2024, 2039, 2061, 2076, 2091, 2108, 2125, 2142, 2159, 2176, 2194, 2212, 2231, 2255, 2274, 2295, 2316, 2332, 2350, 2370, 2385, 2408, 2433, 2458, 2475, 2492, 2509, 2526, 2544, 2562, 2577, 2592, 2609, 2626, 2644, 2661, 2677, 2694, 2712, 2729, 2744, 2760, 2785, 2802, 2818, 2834, 2851, 2869, 2888, 2905, 2922, 2943, 2965, 2983, 3005, 3029, 3046, 3065, 3084, 3101, 3117, 3135, 3161, 3182, 3200, 3219, 3232, 3246, 3260, 3277, 3298, 3316, 3332, 3350, 3377, 3393, 3415, 3438, 3451, 3477, 3504, 3524, 3537, 3561, 3586, 3604, 3627, 3652, 3671, 3692, 3714, 3737, 3754, 3777, 3798, 3826, 3844, 3863, 3881, 3909, 3931, 3964, 3991, 4020, 4043, 4072, 4095, 4129, 4157, 4185, 4213, 4235, 4266, 4285, 4303, 4325, 4342, 4362, 4381, 4407, 4433, 4460, 4491, 4523, 4559, 4587, 4619, 4639, 4663, 4685, 4711, 4724, 4739, 4757, 4777, 4801, 4820, 4842, 4863, 4882, 4903, 4922, 4936, 4964, 4986, 5008, 5034, 5057, 5076, 5094, 5113, 5132, 5153, 5163, 5173, 5188, 5203, 5229, 5255, 5279, 5302, 5321, 5344, 5358, 5376, 5391, 5415, 5435, 5451, 5480, 5499, 5511, 5538, 5557, 5573, 5588, 5605, 5622, 5641, 5658, 5677, 5698, 5717, 5734, 5760, 5784, 5808, 5833, 5855, 5877, 5897, 5919, 5938, 5957, 5975, 5996, 6015, 6030, 6046, 6062, 6078, 6091, 6109, 6128, 6147, 6163, 6187, 6207, 6224, 6242, 6265, 6281, 6299, 6323, 6350, 6373, 6389, 6404, 6420, 6435, 6462, 6488, 6520, 6547, 6575, 6599, 6627, 6650, 6670, 6688, 6706, 6723, 6744, 6761, 6777, 6799, 6819, 6840, 6858, 6874, 6900, 6918, 6940, 6963, 6972, 6987, 7004, 7022, 7048, 7065, 7087, 7108, 7129, 7150, 7169, 7190, 7208, 7226, 7242, 7260, 7279, 7295, 7314, 7341, 7363, 7385, 7406, 7425, 7440, 7456, 7471, 7492, 7507, 7520, 7537, 7562, 7581, 7601, 7620, 7640, 7658, 7686, 7708, 7729, 7755, 7782, 7814, 7843, 7871, 7901, 7930, 7960, 7988, 8017, 8046, 8068, 8098, 8126, 8150, 8176, 8194, 8217, 8245, 8270, 8295, 8320, 8345, 8370, 8391, 8410, 8429, 8443, 8457, 8470, 8491, 8504, 8519, 8541, 8564, 8582, 8600, 8626, 8646, 8674, 8691, 8713, 8731, 8748, 8765, 8787, 8813, 8836, 8856, 8876, 8896, 8910, 8928, 8953, 8975, 8997, 9014, 9043, 9072, 9090, 9116, 9136, 9155, 9176, 9205, 9224, 9252, 9272, 9289, 9297, 9314, 9330, 9346, 9361, 9385, 9400, 9416, 9435, 9449, 9466, 9482, 9496, 9515, 9535, 9558, 9581, 9603, 9612, 9621, 9640, 9664, 9680, 9699, 9717, 9737, 9766, 9776, 9789, 9816, 9833, 9846, 9867, 9896, 9915, 9932, 9950, 9966, 9985, 10006, 10031, 10053, 10068, 10085, 10114, 10136, 10163, 10189, 10201, 10216, 10243, 10271, 10282, 10304, 10328, 10346, 10364, 10382, 10395, 10415, 10435, 10459, 10483, 10506, 10531, 10555, 10579, 10609, 10644, 10657, 10683, 10706, 10729, 10749, 10784, 10812, 10828, 10854, 10880, 10903, 10929, 10955, 10982, 11001, 11024, 11043, 11062, 11075, 11101, 11127, 11142, 11159, 11173, 11192, 11207, 11223, 11240, 11257, 11275, 11289, 11306, 11333, 11352, 11367, 11386, 11405, 11431, 11458, 11481, 11503, 11517, 11538, 11556, 11577, 11598, 11619, 11640, 11669, 11696, 11714, 11733, 11753, 11782, 11802, 11821, 11843, 11869, 11885, 11911, 11928, 11948, 11966, 11986, 12008, 12035, 12052, 12073, 12088, 12110, 12127, 12149, 12173, 12196, 12225, 12247, 12274, 12298, 12328, 12360, 12391, 12406, 12430, 12448, 12473, 12503, 12534, 12562, 12589, 12602, 12627, 12648, 12668, 12688, 12717, 12738, 12759, 12774, 12791, 12811, 12829, 12851, 12871, 12892, 12907, 12929, 12950, 12964, 12983, 13003, 13025, 13052, 13083, 13108, 13132, 13156, 13174, 13198, 13216, 13236, 13261, 13284, 13308, 13332, 13350, 13365, 13398, 13421, 13448, 13465, 13484, 13514, 13537, 13556, 13582, 13601, 13618, 13648, 13670, 13696, 13719, 13742, 13764, 13784, 13802, 13816, 13836, 13850, 13871, 13892, 13917, 13940, 13967, 13983, 14009, 14032, 14051, 14071, 14088, 14115, 14130, 14151, 14170, 14191, 14204, 14229, 14244, 14263, 14282, 14307, 14324, 14342, 14362, 14381, 14409, 14429, 14455, 14480, 14499, 14521, 14544, 14564, 14593, 14617, 14633, 14649, 14670, 14690, 14716, 14739, 14761, 14783, 14809, 14833, 14859, 14874, 14897, 14920, 14941, 14964, 14987, 15013, 15038, 15066, 15098, 15131, 15168, 15195, 15226, 15242, 15264, 15293, 15322, 15350, 15365, 15381, 15397, 15427, 15460, 15480, 15499, 15530, 15559, 15581, 15596, 15612, 15636, 15664, 15687, 15699, 15721, 15740, 15759, 15777, 15802, 15828, 15849, 15869, 15889, 15914, 15944, 15970, 15996, 16027, 16052, 16075, 16104, 16132, 16158, 16186, 16213, 16239, 16266, 16292, 16321, 16351, 16371, 16392, 16421, 16445, 16467, 16488, 16512, 16535, 16560, 16584, 16605, 16626, 16650, 16672, 16692, 16711, 16737, 16762, 16784, 16804, 16818, 16839, 16861, 16878, 16906, 16934, 16960, 16980, 17010, 17037, 17063, 17087, 17109, 17140, 17158, 17184, 17202, 17229, 17263, 17295, 17323, 17345, 17383, 17403, 17427, 17452, 17461, 17510, 17561, 17614, 17639, 17663, 17684, 17707, 17732, 17749, 17765, 17787, 17817, 17868, 17889, 17919, 17941, 17963, 17985, 18007, 18029, 18052, 18073, 18095, 18118, 18140, 18161, 18189, 18210, 18230, 18251, 18264, 18278, 18298, 18319, 18339, 18363, 18387, 18413, 18433, 18454, 18474, 18496, 18517, 18539, 18557, 18575, 18597, 18626, 18659, 18679, 18696, 18728, 18747, 18764, 18778, 18800, 18816, 18839, 18861, 18889, 18913, 18942, 18971, 19002, 19027, 19063, 19093, 19123, 19157, 19192, 19231, 19259, 19292, 19323, 19346, 19374, 19404, 19420, 19436, 19460, 19480, 19505, 19523, 19542, 19563, 19588, 19612, 19639, 19663, 19685, 19702, 19723, 19755, 19796, 19848, 19872, 19904, 19928, 19962, 19997, 20024, 20058, 20091, 20136, 20171, 20193, 20223, 20253, 20279, 20302, 20315, 20338, 20366, 20391, 20416, 20444, 20459, 20478, 20495, 20518, 20548, 20577, 20602, 20634, 20666, 20694, 20725, 20755, 20788, 20805, 20822, 20856, 20879, 20903, 20927, 20943, 20972, 20985, 21005, 21031, 21054, 21071, 21097, 21121, 21149, 21175, 21191, 21214, 21239, 21268, 21297, 21322, 21343, 21365, 21388, 21411, 21439, 21460, 21481, 21499, 21519, 21543, 21568, 21590, 21614, 21636, 21655, 21681, 21707, 21732, 21752, 21778, 21801, 21827, 21856, 21886, 21910, 21932, 21949, 21969, 21989, 22013, 22032, 22052, 22075, 22096, 22111, 22125, 22150, 22172, 22194, 22215, 22245, 22270, 22297, 22322, 22343, 22368, 22401, 22438, 22475, 22490, 22517, 22537, 22562, 22586, 22611, 22627, 22646, 22672, 22699, 22725, 22744, 22762, 22782, 22807, 22829, 22866, 22897, 22927, 22964, 22991, 23018, 23039, 23063, 23087, 23110, 23137, 23164, 23193, 23211, 23244, 23265, 23288, 23311, 23337, 23360, 23386, 23410, 23433, 23459, 23485, 23509, 23535, 23559, 23592, 23618, 23654, 23681, 23705, 23731, 23755, 23781, 23815, 23841, 23869, 23901, 23935, 23967, 24007, 24046, 24079, 24112, 24154, 24176, 24205, 24235, 24250, 24267, 24282, 24297, 24313, 24328, 24343, 24358, 24374, 24399, 24420, 24437, 24461, 24482, 24505, 24524, 24546, 24575, 24598, 24626, 24658, 24684, 24716, 24748, 24786, 24816, 24849, 24867, 24886, 24905, 24919, 24942, 24971, 24991, 25022, 25055, 25086, 25117, 25146, 25169, 25203, 25226, 25245, 25264, 25283, 25308, 25338, 25368, 25400, 25421, 25445, 25465, 25490, 25510, 25528, 25545, 25568, 25594, 25612, 25630, 25647, 25664, 25681, 25700, 25716, 25733, 25750, 25772, 25792, 25823, 25854, 25885, 25905, 25933, 25948, 25966, 25986, 26002, 26024, 26053, 26075, 26097, 26136, 26161, 26202, 26223, 26247, 26269, 26294, 26319, 26346, 26361, 26386, 26418, 26442, 26467, 26498, 26512, 26536, 26564, 26599, 26631, 26663, 26697, 26731, 26761, 26782, 26798, 26821, 26851, 26890, 26902, 26914, 26940, 26974, 26997, 27020, 27051, 27079, 27096, 27122, 27152, 27189, 27214, 27240, 27271, 27296, 27332, 27356, 27381, 27409, 27435, 27456, 27481, 27510, 27534, 27562, 27583, 27615, 27646, 27668, 27697, 27712, 27727, 27751, 27775, 27803, 27823, 27850, 27883, 27925, 27953, 27986, 28028, 28062, 28105, 28139, 28167, 28200, 28233, 28271, 28320, 28343, 28362, 28381, 28405, 28429, 28452, 28481, 28510, 28539, 28554, 28569, 28601, 28633, 28654, 28674, 28706, 28727, 28748, 28770, 28792, 28825, 28852, 28888, 28906, 28935, 28965, 28992, 29019, 29041, 29071, 29101, 29118, 29141, 29165, 29185, 29208, 29226, 29248, 29278, 29305, 29329, 29353, 29381, 29408, 29439, 29465, 29494, 29524, 29563, 29602, 29641, 29660, 29673}

func (i WorldType) String() string {
	i -= 1
	if i >= WorldType(len(_WorldType_index)-1) {
		return fmt.Sprintf("WorldType(%d)", i+1)
	}
	return _WorldType_name[_WorldType_index[i]:_WorldType_index[i+1]]
}