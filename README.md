# Satlantis Database Models

These are defined to be used with GORM. Functions which interact with these models are kept in the repositories which use these models. This aligns with the principle of separation of concerns and will lead to a more modular codebase.

### Separation of Concerns:

Model definitions are separated from the database interaction logic, leading to a cleaner architecture.

### Reusability:

Models can be reused across multiple projects without duplicating code.

### Independent Versioning:

Models can be versioned independently, ensuring that changes in models do not directly affect the repositories using them until explicitly updated.

### Flexibility:

Each repository can define its own database interaction logic, allowing for flexibility in handling different use cases or database setups.
