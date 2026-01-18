# Tardigrade-Mod Project Design

## Overview

Tardigrade-Mod is a lightweight, file-based NoSQL database library designed specifically for small to medium-sized Go applications. It provides a simple key-value storage mechanism with JSON serialization, making it ideal for applications that need persistent storage without the overhead of a full database server.

## Architecture

### Core Design Philosophy

- **Zero Dependencies**: Pure Go implementation with only standard library dependencies
- **File-Based Storage**: Single-file database approach for simplicity and portability
- **JSON-Native**: All data stored and retrieved in JSON format
- **Multi-Database Support**: Applications can manage multiple database files simultaneously
- **Cross-Platform**: Compatible with Linux, macOS, and Windows

### System Architecture

```
┌─────────────────────────────────────────────────────────┐
│                   Application Layer                      │
│              (Your Go Application)                       │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│              Tardigrade API Layer                        │
│  ┌──────────────┬──────────────┬──────────────────────┐ │
│  │   CRUD Ops   │  Query Ops   │   Utility Ops        │ │
│  │  - AddField  │ - SelectByID │  - CreateDB          │ │
│  │  - Modify    │ - Search     │  - DeleteDB          │ │
│  │  - Remove    │ - FirstX     │  - EmptyDB           │ │
│  │              │ - LastX      │  - CreatedDBCopy     │ │
│  └──────────────┴──────────────┴──────────────────────┘ │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│              Data Processing Layer                       │
│  ┌──────────────┬──────────────┬──────────────────────┐ │
│  │ Serialization│  Encryption  │   Encoding           │ │
│  │  - MyMarshal │ - MyEncrypt  │  - MyEncode          │ │
│  │  - MyIndent  │ - MyDecrypt  │  - MyDecode          │ │
│  └──────────────┴──────────────┴──────────────────────┘ │
└────────────────────┬────────────────────────────────────┘
                     │
                     ▼
┌─────────────────────────────────────────────────────────┐
│              File System Layer                           │
│         (OS-Specific File Operations)                    │
└─────────────────────────────────────────────────────────┘
```

## Data Model

### Record Structure

Each record in the database follows a consistent JSON schema:

```json
{
  "id": 1,
  "key": "record_identifier",
  "data": "actual_data_content"
}
```

**Fields:**
- `id` (int): Auto-incrementing unique identifier
- `key` (string): User-defined key/label for the record
- `data` (string): The actual data payload (can be any string, including serialized JSON)

### Storage Format

- **File Format**: Newline-delimited JSON (NDJSON)
- **Encoding**: UTF-8
- **Line Terminator**: `\n` (Unix-style)
- **File Extension**: `.db` (convention, not enforced)

## Core Capabilities

### 1. Database Management

#### Create Database
- Creates new database file if it doesn't exist
- Returns status and full path
- Automatic directory resolution

#### Delete Database
- Permanently removes database file
- Safety checks before deletion
- Returns confirmation status

#### Empty Database
- Clears all records while preserving file
- Atomic operation (delete + recreate)
- Maintains file permissions

#### Backup Database
- Creates copy in user's home directory
- Preserves all data integrity
- Buffer-based copying for efficiency

### 2. CRUD Operations

#### Create (AddField)
- Auto-incrementing ID generation
- Atomic append operations
- Automatic database creation if missing
- Returns boolean success status

#### Read Operations
- **SelectByID**: Direct ID-based lookup with O(n) complexity
- **FirstField**: Retrieve first record
- **LastField**: Retrieve last record
- **FirstXFields**: Batch retrieval from start
- **LastXFields**: Batch retrieval from end
- **SelectSearch**: Multi-pattern search with AND logic

#### Update (ModifyField)
- In-place record modification
- Preserves ID and record order
- Atomic file rewrite operation

#### Delete (RemoveField)
- Removes specific record by ID
- Maintains file integrity
- Returns deleted record for confirmation

### 3. Query Capabilities

#### Output Formats
All read operations support multiple output formats:
- **raw**: Single-line JSON string
- **json**: Pretty-printed JSON with indentation
- **id**: Only the ID field
- **key**: Only the key field
- **value**: Only the data field

#### Search Functionality
- Multi-keyword search (comma or space-separated)
- Case-insensitive matching
- AND logic (all keywords must match)
- Searches across all fields (id, key, data)

### 4. Utility Functions

#### Data Serialization
- **MyMarshal**: JSON encoding with HTML escape disabled
- **MyIndent**: Pretty-print JSON with custom formatting

#### Encoding/Decoding
- **MyEncode**: Base64 encoding
- **MyDecode**: Base64 decoding

#### Encryption (AES-CFB)
- **MyEncrypt**: AES encryption with custom password
- **MyDecrypt**: Decryption with password verification
- 16-byte initialization vector
- Suitable for sensitive data storage

#### Metadata
- **GetVersion**: Returns current release version
- **GetUpdated**: Returns last update timestamp
- **CountSize**: Returns total record count
- **UniqueID**: Returns last used ID for auto-increment

### 5. Cross-Platform Support

#### OS Detection
- Automatic path separator detection
- Windows: `\`
- Linux/macOS: `/`
- Ensures file path compatibility

## Technical Specifications

### Performance Characteristics

| Operation | Time Complexity | Space Complexity | Notes |
|-----------|----------------|------------------|-------|
| AddField | O(1) | O(1) | Append-only operation |
| SelectByID | O(n) | O(1) | Linear scan required |
| RemoveField | O(n) | O(n) | Full file rewrite |
| ModifyField | O(n) | O(n) | Full file rewrite |
| CountSize | O(n) | O(1) | Byte-level scanning |
| FirstField | O(1) | O(1) | Reads first line only |
| LastField | O(n) | O(1) | Scans to last line |
| SelectSearch | O(n*m) | O(k) | n=records, m=keywords, k=results |

### Scalability Considerations

**Optimal Use Cases:**
- Record count: < 100,000 records
- File size: < 100 MB
- Concurrent readers: Multiple (read-only)
- Concurrent writers: Single writer recommended

**Limitations:**
- No indexing (linear search for all queries)
- No transaction support
- No concurrent write safety
- Full file rewrite for updates/deletes
- Memory-intensive for large result sets

### Error Handling

- Panic-based error handling via `CheckError` function
- Errors logged before panic
- Common error scenarios:
  - File not found
  - Permission denied
  - Buffer overflow
  - JSON parsing errors
  - Encryption/decryption failures

## Security Features

### Encryption Support

- **Algorithm**: AES (Advanced Encryption Standard)
- **Mode**: CFB (Cipher Feedback)
- **Key Size**: 128/192/256-bit (password-dependent)
- **Encoding**: Base64 for encrypted output

### Security Considerations

- Passwords must be exactly 16, 24, or 32 bytes for AES
- No built-in password hashing (user responsibility)
- Encryption is opt-in (not automatic)
- No field-level encryption in standard operations

## Use Cases

### Ideal Scenarios

1. **Configuration Storage**: Application settings and preferences
2. **Cache Layer**: Temporary data storage with persistence
3. **Session Management**: User session data for web applications
4. **Logging**: Structured log storage with search capabilities
5. **Prototyping**: Quick data persistence during development
6. **Embedded Systems**: Lightweight storage for IoT devices
7. **CLI Tools**: Data storage for command-line applications

### Not Recommended For

1. High-concurrency applications (>10 concurrent writers)
2. Large datasets (>1M records)
3. Real-time analytics
4. Applications requiring ACID transactions
5. Multi-user systems without external locking

## Future Enhancements

### Planned Features (Outstanding)

```go
// Encrypted field operations
func (*Tardigrade).AddCryptField(key string, data string, db string) bool
func (*Tardigrade).SelectByIDdecrypt(id int, f string, db string) string
```

### Potential Improvements

1. **Indexing**: B-tree or hash-based indexing for faster lookups
2. **Compression**: Optional gzip compression for storage efficiency
3. **Transactions**: Basic transaction support with rollback
4. **Locking**: File-based locking for concurrent write safety
5. **Streaming**: Iterator-based API for large datasets
6. **Schema Validation**: Optional JSON schema validation
7. **Backup Rotation**: Automatic backup with retention policies
8. **Query Language**: Simple query DSL for complex searches

## Integration Guide

### Installation

```bash
go get github.com/gcclinux/tardigrade-mod
```

### Basic Usage Pattern

```go
import "github.com/gcclinux/tardigrade-mod"

// Initialize
tar := tardigrade.Tardigrade{}

// Create database
tar.CreateDB("myapp.db")

// Add data
tar.AddField("user:1", "John Doe", "myapp.db")

// Query data
result := tar.SelectByID(1, "json", "myapp.db")

// Search
format, results := tar.SelectSearch("John", "json", "myapp.db")
```

### Best Practices

1. **Database Naming**: Use descriptive names with `.db` extension
2. **Error Handling**: Wrap operations in defer/recover for panic handling
3. **Backups**: Regular backups using `CreatedDBCopy`
4. **Data Validation**: Validate data before `AddField`
5. **Search Optimization**: Use specific keywords for faster searches
6. **Batch Operations**: Use `FirstXFields`/`LastXFields` for bulk reads
7. **Encryption**: Encrypt sensitive data before storage

## Version History

- **0.0.1**: Initial version
- **0.0.2**: README updates and bug fixes
- **0.0.3**: Converted to structure method pattern
- **0.0.4**: Renamed from tardigrade app to tardigrade-mod
- **0.1.0**: Added multiple functions from original app
- **0.1.2**: Fixed lowercase string return bug
- **0.1.3**: Fixed lowercase function accessibility
- **0.1.4**: Fixed HTML escape in JSON encoding
- **0.2.0**: Added utility functions to main struct
- **0.2.1**: Minor bug fixes
- **0.2.3**: Work in progress on data encryption
- **0.2.5**: Multi-database support (current)

## License

See LICENSE file for details.

## Contributing

Contributions welcome! Please ensure:
- Code follows Go conventions
- All functions include error handling
- Documentation is updated
- Examples are provided for new features
