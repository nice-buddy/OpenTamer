#ifndef OPENTAMER_L10N_BRIDGE_H
#define OPENTAMER_L10N_BRIDGE_H

// Forward declarations for l10n functions exported from Go
// These are defined in l10n_bridge_darwin.go with //export directives
char *OpenTamerL10nTranslate(char *language, char *key);
char *OpenTamerL10nStatus(char *language, char *status);
char *OpenTamerL10nSafety(char *language, char *reason);
void OpenTamerL10nFree(char *value);

#endif
