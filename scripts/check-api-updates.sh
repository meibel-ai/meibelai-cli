#!/bin/bash
# Script to check for API specification updates

set -e

CURRENT_SPEC="api.json"
NEW_SPEC="api-new.json"
SPEC_URL="https://storage.googleapis.com/meibel-api-docs/api.json"
CHANGES_FILE="api-changes.md"

echo "Checking for API updates..."

# Download latest spec
echo "Downloading latest API specification..."
if ! curl -s "$SPEC_URL" > "$NEW_SPEC"; then
    echo "Error: Failed to download API specification"
    exit 1
fi

# Check if current spec exists
if [ ! -f "$CURRENT_SPEC" ]; then
    echo "No current spec found. Saving new spec as current."
    mv "$NEW_SPEC" "$CURRENT_SPEC"
    exit 0
fi

# Check if spec has changed
if ! diff -q "$CURRENT_SPEC" "$NEW_SPEC" > /dev/null 2>&1; then
    echo "API specification has changed!"
    
    # Generate changelog
    echo "## API Changes" > "$CHANGES_FILE"
    echo "" >> "$CHANGES_FILE"
    echo "Date: $(date)" >> "$CHANGES_FILE"
    echo "" >> "$CHANGES_FILE"
    
    # Check if jq is installed
    if command -v jq &> /dev/null; then
        echo "### Endpoint Changes:" >> "$CHANGES_FILE"
        echo "" >> "$CHANGES_FILE"
        
        # Extract endpoints
        jq -r '.paths | keys[]' "$NEW_SPEC" | sort > new-endpoints.txt
        jq -r '.paths | keys[]' "$CURRENT_SPEC" | sort > old-endpoints.txt
        
        # Find new endpoints
        NEW_ENDPOINTS=$(comm -13 old-endpoints.txt new-endpoints.txt)
        if [ -n "$NEW_ENDPOINTS" ]; then
            echo "#### New Endpoints:" >> "$CHANGES_FILE"
            echo "$NEW_ENDPOINTS" | while read -r endpoint; do
                echo "- $endpoint" >> "$CHANGES_FILE"
            done
            echo "" >> "$CHANGES_FILE"
        fi
        
        # Find removed endpoints
        REMOVED_ENDPOINTS=$(comm -23 old-endpoints.txt new-endpoints.txt)
        if [ -n "$REMOVED_ENDPOINTS" ]; then
            echo "#### Removed Endpoints:" >> "$CHANGES_FILE"
            echo "$REMOVED_ENDPOINTS" | while read -r endpoint; do
                echo "- $endpoint" >> "$CHANGES_FILE"
            done
            echo "" >> "$CHANGES_FILE"
        fi
        
        # Check version change
        OLD_VERSION=$(jq -r '.info.version' "$CURRENT_SPEC")
        NEW_VERSION=$(jq -r '.info.version' "$NEW_SPEC")
        
        if [ "$OLD_VERSION" != "$NEW_VERSION" ]; then
            echo "### Version Change:" >> "$CHANGES_FILE"
            echo "- Previous: $OLD_VERSION" >> "$CHANGES_FILE"
            echo "- Current: $NEW_VERSION" >> "$CHANGES_FILE"
            echo "" >> "$CHANGES_FILE"
        fi
        
        # Clean up temp files
        rm -f old-endpoints.txt new-endpoints.txt
    else
        echo "Warning: jq not installed. Install it for detailed change analysis." >> "$CHANGES_FILE"
        echo "" >> "$CHANGES_FILE"
        echo "To install jq:" >> "$CHANGES_FILE"
        echo "- macOS: brew install jq" >> "$CHANGES_FILE"
        echo "- Ubuntu: apt-get install jq" >> "$CHANGES_FILE"
        echo "" >> "$CHANGES_FILE"
    fi
    
    # Show basic diff stats
    echo "### Change Statistics:" >> "$CHANGES_FILE"
    diff -u "$CURRENT_SPEC" "$NEW_SPEC" | diffstat -p1 >> "$CHANGES_FILE" 2>/dev/null || true
    
    # Update the spec
    mv "$NEW_SPEC" "$CURRENT_SPEC"
    
    echo ""
    echo "Changes detected and saved to $CHANGES_FILE"
    echo "Please review the changes and update the CLI accordingly."
    
    # Exit with code 1 to indicate changes were found
    exit 1
else
    echo "No API changes detected"
    rm -f "$NEW_SPEC"
    exit 0
fi