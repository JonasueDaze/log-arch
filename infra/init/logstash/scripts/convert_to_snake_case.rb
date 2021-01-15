def register(params)
end

def filter(event)
    hash = convert_hash_keys(event.to_hash)
    hash.each do |k,v|
        event.set(k.to_s, v)
    end
    return [event]
end

def convert_hash_keys(value)
    case value
        when Array
            value.map { |v| convert_hash_keys(v) }
        when Hash
            Hash[value.map { |k, v| [underscore_key(k), convert_hash_keys(v)] }]
        else
            value
    end
end

def underscore_key(k)
    to_snake_case(k.to_s).to_sym
end

def to_snake_case(string)
    string.gsub(/::/, '/').
    gsub(/([A-Z]+)([A-Z][a-z])/,'\1_\2').
    gsub(/([a-z\d])([A-Z])/,'\1_\2').
    tr("-", "_").
    downcase
end
