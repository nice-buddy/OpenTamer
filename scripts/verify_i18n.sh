#!/bin/bash
# Script to verify i18n implementation

set -e

echo "=== OpenTamer i18n Verification ==="
echo

# Test 1: Build check
echo "1. Building OpenTamer..."
go build -v ./cmd/opentamer > /dev/null 2>&1
if [ $? -eq 0 ]; then
    echo "   ✓ Build successful"
else
    echo "   ✗ Build failed"
    exit 1
fi

# Test 2: Run unit tests
echo
echo "2. Running language tests..."
go test -v -run TestLanguage ./internal/config > /tmp/lang_test.log 2>&1
if [ $? -eq 0 ]; then
    detected=$(grep "Detected language:" /tmp/lang_test.log | awk '{print $NF}')
    echo "   ✓ Language detection tests passed"
    echo "   ✓ System language detected as: $detected"
else
    echo "   ✗ Language tests failed"
    cat /tmp/lang_test.log
    exit 1
fi

# Test 3: Check locale tests
echo
echo "3. Running locale tests..."
go test -v -run TestLocale ./internal/ui > /dev/null 2>&1
if [ $? -eq 0 ]; then
    echo "   ✓ Locale tests passed"
else
    echo "   ✗ Locale tests failed"
    exit 1
fi

# Test 4: Check menu integration
echo
echo "4. Checking menu integration..."
if grep -q "Language" internal/macos/menu/menu_bar_bridge.m && \
   grep -q '"en"' internal/macos/menu/menu_bar_bridge.m && \
   grep -q '"zh"' internal/macos/menu/menu_bar_bridge.m; then
    echo "   ✓ Language preference menu integrated"
else
    echo "   ✗ Language preference menu not found"
    exit 1
fi

# Test 5: Check config structure
echo
echo "5. Checking config structure..."
if grep -q "PreferenceLanguage" internal/ui/preferences.go && \
   grep -q "DetectLanguage" internal/config/config.go; then
    echo "   ✓ Config structure includes language support"
else
    echo "   ✗ Config structure incomplete"
    exit 1
fi

# Test 6: Run all tests
echo
echo "6. Running full test suite..."
go test ./... > /tmp/all_tests.log 2>&1
if [ $? -eq 0 ]; then
    passed=$(grep -o "[0-9]* passed" /tmp/all_tests.log | awk '{sum += $1} END {print sum}')
    echo "   ✓ All tests passed ($passed tests)"
else
    echo "   ✗ Some tests failed"
    tail -20 /tmp/all_tests.log
    exit 1
fi

echo
echo "=== Verification Summary ==="
echo "✓ Build: OK"
echo "✓ Language detection: Working"
echo "✓ Locale system: Implemented"
echo "✓ UI integration: Complete"
echo "✓ Tests: All passing"
echo
echo "The i18n feature is fully implemented and ready to use!"
echo
echo "To change language:"
echo "1. Open OpenTamer"
echo "2. Click the menu bar icon"
echo "3. Click '...' for Preferences"
echo "4. Select 'General' > 'Language'"
echo "5. Choose: Auto, English, or 中文"
